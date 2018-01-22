package wallet

import (
	"fmt"
	"os"
	"sync"

	"github.com/skycoin/skycoin/src/cipher"
	bip39 "github.com/skycoin/skycoin/src/cipher/go-bip39"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/visor/blockdb"
)

// BalanceGetter interface for getting the balance of given addresses
type BalanceGetter interface {
	GetBalanceOfAddrs(addrs []cipher.Address) ([]BalancePair, error)
}

// Service wallet service struct
type Service struct {
	sync.RWMutex
	wallets          Wallets
	firstAddrIDMap   map[string]string // Key: first address in wallet; Value: wallet id
	walletDirectory  string
	disableWalletAPI bool
}

// NewService new wallet service
func NewService(walletDir string, disableWalletAPI bool) (*Service, error) {
	serv := &Service{
		firstAddrIDMap:   make(map[string]string),
		disableWalletAPI: disableWalletAPI,
	}

	if serv.disableWalletAPI {
		return serv, nil
	}

	if err := os.MkdirAll(walletDir, os.FileMode(0700)); err != nil {
		return nil, fmt.Errorf("failed to create wallet directory %s: %v", walletDir, err)
	}

	serv.walletDirectory = walletDir

	// Removes .wlt.bak files before loading wallets
	if err := removeBackupFiles(serv.walletDirectory); err != nil {
		return nil, fmt.Errorf("remove .wlt.bak files in %v failed: %v", serv.walletDirectory, err)
	}

	// Loads wallets
	w, err := LoadWallets(serv.walletDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to load all wallets: %v", err)
	}

	serv.wallets = serv.removeDup(w)

	if len(serv.wallets) == 0 {
		seed, err := bip39.NewDefaultMnemomic()
		if err != nil {
			return nil, err
		}

		// Create default wallet
		w, err := serv.CreateWallet("", Options{
			Label:      "Your Wallet",
			Seed:       seed,
			CryptoType: DefaultCryptoType,
		})
		if err != nil {
			return nil, err
		}

		if err := Save(serv.walletDirectory, w); err != nil {
			return nil, fmt.Errorf("failed to save wallets to %s: %v", serv.walletDirectory, err)
		}
	}

	return serv, nil
}

// CreateWallet creates a wallet with one address
func (serv *Service) CreateWallet(wltName string, options Options) (*Wallet, error) {
	serv.Lock()
	defer serv.Unlock()
	if serv.disableWalletAPI {
		return nil, ErrWalletApiDisabled
	}
	if wltName == "" {
		wltName = serv.generateUniqueWalletFilename()
	}

	return serv.loadWallet(wltName, options, 0, nil)
}

// ScanAheadWalletAddresses scans n addresses for a balance, and sets the wallet's entry list to the highest
// address with a non-zero coins balance.
// Set password as nil if the wallet is not encrypted, otherwise the password must be provided.
func (serv *Service) ScanAheadWalletAddresses(wltName string, password []byte, scanN uint64, bg BalanceGetter) (*Wallet, error) {
	serv.Lock()
	defer serv.Unlock()

	w, err := serv.getWallet(wltName)
	if err != nil {
		return nil, err
	}

	f := func(wlt *Wallet) error {
		return wlt.ScanAddresses(scanN, bg)
	}

	if w.IsEncrypted() {
		if err := w.guard(password, f); err != nil {
			return nil, err
		}
	} else {
		if err := f(w); err != nil {
			return nil, err
		}
	}

	if err := Save(serv.walletDirectory, w); err != nil {
		return nil, err
	}

	serv.wallets.set(w)

	return w.clone(), nil
}

// loadWallet loads wallet from seed and scan the first N addresses
func (serv *Service) loadWallet(wltName string, options Options, scanN uint64, bg BalanceGetter) (*Wallet, error) {
	// Creates the wallet
	w, err := NewWallet(wltName, options)
	if err != nil {
		return nil, err
	}

	f := func(wlt *Wallet) error {
		if len(wlt.Entries) == 0 {
			// Generate a default address
			wlt.GenerateAddresses(1)
		}

		// Check for duplicate wallets by initial seed
		if id, ok := serv.firstAddrIDMap[wlt.Entries[0].Address.String()]; ok {
			return fmt.Errorf("wallet %s would be duplicate with %v, same seed", wlt.Filename(), id)
		}

		// Scan for addresses with balances
		if scanN > 1 && bg != nil {
			if err := wlt.ScanAddresses(scanN-1, bg); err != nil {
				return err
			}
		}
		return nil
	}

	if w.IsEncrypted() {
		if err := w.guard(options.Password, f); err != nil {
			return nil, err
		}
	} else {
		if err := f(w); err != nil {
			return nil, err
		}
	}

	if err := serv.wallets.add(w); err != nil {
		return nil, err
	}

	if err := Save(serv.walletDirectory, w); err != nil {
		// If save fails, remove the added wallet
		serv.wallets.remove(w.Filename())
		return nil, err
	}

	serv.firstAddrIDMap[w.Entries[0].Address.String()] = w.Filename()

	return w.clone(), nil
}

func (serv *Service) generateUniqueWalletFilename() string {
	wltName := newWalletFilename()
	for {
		if _, ok := serv.wallets.get(wltName); !ok {
			break
		}
		wltName = newWalletFilename()
	}

	return wltName
}

// EncryptWallet encrypts wallet with password
func (serv *Service) EncryptWallet(wltID string, password []byte) error {
	serv.Lock()
	defer serv.Unlock()
	w, err := serv.getWallet(wltID)
	if err != nil {
		return err
	}

	if w.IsEncrypted() {
		return ErrWalletEncrypted
	}

	if err := w.lock(password); err != nil {
		return err
	}

	// Save to disk first
	if err := Save(serv.walletDirectory, w); err != nil {
		return err
	}

	// Sets the encrypted wallet
	serv.wallets.set(w)
	return nil
}

// DecryptWallet decrypts wallet with password
func (serv *Service) DecryptWallet(wltID string, password []byte) error {
	serv.Lock()
	defer serv.Unlock()
	w, err := serv.getWallet(wltID)
	if err != nil {
		return err
	}

	// Returns error if wallet is not encrypted
	if !w.IsEncrypted() {
		return ErrWalletNotEncrypted
	}

	// Unlocks the wallet
	unlockWlt, err := w.unlock(password)
	if err != nil {
		return err
	}

	// Updates the wallet file
	if err := Save(serv.walletDirectory, unlockWlt); err != nil {
		return err
	}

	// Sets the decrypted wallet in memory
	serv.wallets.set(unlockWlt)
	return nil
}

// NewAddresses generate address entries in given wallet,
// return nil if wallet does not exist.
// Set password as nil if the wallet is not encrypted, otherwise the password must be provided.
func (serv *Service) NewAddresses(wltID string, password []byte, num uint64) ([]cipher.Address, error) {
	serv.Lock()
	defer serv.Unlock()
	w, err := serv.getWallet(wltID)
	if err != nil {
		return nil, err
	}

	var addrs []cipher.Address
	f := func(wlt *Wallet) error {
		var err error
		addrs, err = wlt.GenerateAddresses(num)
		return err
	}

	if w.IsEncrypted() {
		if err := w.guard(password, f); err != nil {
			return nil, err
		}
	} else {
		if err := f(w); err != nil {
			return nil, err
		}
	}

	// Set the updated wallet back
	serv.wallets.set(w)

	if err := Save(serv.walletDirectory, w); err != nil {
		return []cipher.Address{}, err
	}

	return addrs, nil
}

// GetAddresses returns all addresses in given wallet
func (serv *Service) GetAddresses(wltID string) ([]cipher.Address, error) {
	serv.RLock()
	defer serv.RUnlock()
	w, err := serv.getWallet(wltID)
	if err != nil {
		return nil, err
	}

	return w.GetAddresses(), nil
}

// GetWallet returns wallet by id
func (serv *Service) GetWallet(wltID string) (*Wallet, error) {
	serv.RLock()
	defer serv.RUnlock()

	return serv.getWallet(wltID)
}

// returns the clone of the wallet of given id
func (serv *Service) getWallet(wltID string) (*Wallet, error) {
	w, ok := serv.wallets.get(wltID)
	if !ok {
		return nil, ErrWalletNotExist
	}
	return w.clone(), nil
}

// GetWallets returns all wallet clones
func (serv *Service) GetWallets() Wallets {
	serv.RLock()
	defer serv.RUnlock()
	wlts := make(Wallets, len(serv.wallets))
	for k, w := range serv.wallets {
		wlts[k] = w.clone()
	}
	return wlts
}

// ReloadWallets reload wallets
func (serv *Service) ReloadWallets() error {
	serv.Lock()
	defer serv.Unlock()
	if serv.disableWalletAPI {
		return ErrWalletApiDisabled
	}
	wallets, err := LoadWallets(serv.walletDirectory)
	if err != nil {
		return err
	}

	serv.firstAddrIDMap = make(map[string]string)
	serv.wallets = serv.removeDup(wallets)
	return nil
}

// CreateAndSignTransaction creates and sign transaction from wallet
// Set password as nil if the wallet is not encrypted, otherwise the password must be provided
func (serv *Service) CreateAndSignTransaction(wltID string, password []byte, vld Validator, unspent blockdb.UnspentGetter,
	headTime, coins uint64, dest cipher.Address) (*coin.Transaction, error) {
	serv.RLock()
	defer serv.RUnlock()
	w, err := serv.getWallet(wltID)
	if err != nil {
		return nil, err
	}

	var tx *coin.Transaction
	f := func(wlt *Wallet) error {
		var err error
		tx, err = wlt.CreateAndSignTransaction(vld, unspent, headTime, coins, dest)
		return err
	}

	if w.IsEncrypted() {
		dw, err := w.unlock(password)
		if err != nil {
			return nil, err
		}

		if err := f(dw); err != nil {
			return nil, err
		}
	} else {
		if err := f(w); err != nil {
			return nil, err
		}
	}
	return tx, nil
}

// UpdateWalletLabel updates the wallet label
func (serv *Service) UpdateWalletLabel(wltID, label string) error {
	serv.Lock()
	defer serv.Unlock()
	var wlt *Wallet
	if err := serv.wallets.update(wltID, func(w *Wallet) (*Wallet, error) {
		w.setLabel(label)
		wlt = w
		return w, nil
	}); err != nil {
		return err
	}

	return Save(serv.walletDirectory, wlt)
}

// Remove removes wallet of given wallet id from the service
func (serv *Service) Remove(wltID string) {
	serv.Lock()
	defer serv.Unlock()
	serv.wallets.Remove(wltID)
}

func (serv *Service) removeDup(wlts Wallets) Wallets {
	var rmWltIDS []string
	// remove dup wallets
	for wltID, wlt := range wlts {
		if len(wlt.Entries) == 0 {
			// empty wallet
			rmWltIDS = append(rmWltIDS, wltID)
			continue
		}

		addr := wlt.Entries[0].Address.String()
		id, ok := serv.firstAddrIDMap[addr]

		if ok {
			// check whose entries number is bigger
			pw, _ := wlts.get(id)

			if len(pw.Entries) >= len(wlt.Entries) {
				rmWltIDS = append(rmWltIDS, wltID)
				continue
			}

			// replace the old wallet with the new one
			// records the wallet id that need to remove
			rmWltIDS = append(rmWltIDS, id)
			// update wallet id
			serv.firstAddrIDMap[addr] = wltID
			continue
		}

		serv.firstAddrIDMap[addr] = wltID
	}

	// remove the duplicate and empty wallet
	for _, id := range rmWltIDS {
		wlts.remove(id)
	}

	return wlts
}
