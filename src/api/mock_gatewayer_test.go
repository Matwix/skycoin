// Code generated by mockery v1.0.0. DO NOT EDIT.

package api

import (
	cipher "github.com/skycoin/skycoin/src/cipher"
	coin "github.com/skycoin/skycoin/src/coin"

	daemon "github.com/skycoin/skycoin/src/daemon"

	historydb "github.com/skycoin/skycoin/src/visor/historydb"

	kvstorage "github.com/skycoin/skycoin/src/kvstorage"

	mock "github.com/stretchr/testify/mock"

	time "time"

	transaction "github.com/skycoin/skycoin/src/transaction"

	visor "github.com/skycoin/skycoin/src/visor"

	wallet "github.com/skycoin/skycoin/src/wallet"
)

// MockGatewayer is an autogenerated mock type for the Gatewayer type
type MockGatewayer struct {
	mock.Mock
}

// AddStorageValue provides a mock function with given fields: storageType, key, val
func (_m *MockGatewayer) AddStorageValue(storageType kvstorage.Type, key string, val string) error {
	ret := _m.Called(storageType, key, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(kvstorage.Type, string, string) error); ok {
		r0 = rf(storageType, key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddressCount provides a mock function with given fields:
func (_m *MockGatewayer) AddressCount() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddressesActivity provides a mock function with given fields: addrs
func (_m *MockGatewayer) AddressesActivity(addrs []cipher.Addresser) ([]bool, error) {
	ret := _m.Called(addrs)

	var r0 []bool
	if rf, ok := ret.Get(0).(func([]cipher.Addresser) []bool); ok {
		r0 = rf(addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]cipher.Addresser) error); ok {
		r1 = rf(addrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransaction provides a mock function with given fields: p, wp
func (_m *MockGatewayer) CreateTransaction(p transaction.Params, wp visor.CreateTransactionParams) (*coin.Transaction, []visor.TransactionInput, error) {
	ret := _m.Called(p, wp)

	var r0 *coin.Transaction
	if rf, ok := ret.Get(0).(func(transaction.Params, visor.CreateTransactionParams) *coin.Transaction); ok {
		r0 = rf(p, wp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.Transaction)
		}
	}

	var r1 []visor.TransactionInput
	if rf, ok := ret.Get(1).(func(transaction.Params, visor.CreateTransactionParams) []visor.TransactionInput); ok {
		r1 = rf(p, wp)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(transaction.Params, visor.CreateTransactionParams) error); ok {
		r2 = rf(p, wp)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateWallet provides a mock function with given fields: wltName, options
func (_m *MockGatewayer) CreateWallet(wltName string, options wallet.Options) (wallet.Wallet, error) {
	ret := _m.Called(wltName, options)

	var r0 wallet.Wallet
	if rf, ok := ret.Get(0).(func(string, wallet.Options) wallet.Wallet); ok {
		r0 = rf(wltName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, wallet.Options) error); ok {
		r1 = rf(wltName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DaemonConfig provides a mock function with given fields:
func (_m *MockGatewayer) DaemonConfig() daemon.DaemonConfig {
	ret := _m.Called()

	var r0 daemon.DaemonConfig
	if rf, ok := ret.Get(0).(func() daemon.DaemonConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(daemon.DaemonConfig)
	}

	return r0
}

// DecryptWallet provides a mock function with given fields: wltID, password
func (_m *MockGatewayer) DecryptWallet(wltID string, password []byte) (wallet.Wallet, error) {
	ret := _m.Called(wltID, password)

	var r0 wallet.Wallet
	if rf, ok := ret.Get(0).(func(string, []byte) wallet.Wallet); ok {
		r0 = rf(wltID, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(wltID, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisconnectByGnetID provides a mock function with given fields: gnetID
func (_m *MockGatewayer) DisconnectByGnetID(gnetID uint64) error {
	ret := _m.Called(gnetID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(gnetID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EncryptWallet provides a mock function with given fields: wltID, password
func (_m *MockGatewayer) EncryptWallet(wltID string, password []byte) (wallet.Wallet, error) {
	ret := _m.Called(wltID, password)

	var r0 wallet.Wallet
	if rf, ok := ret.Get(0).(func(string, []byte) wallet.Wallet); ok {
		r0 = rf(wltID, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(wltID, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllStorageValues provides a mock function with given fields: storageType
func (_m *MockGatewayer) GetAllStorageValues(storageType kvstorage.Type) (map[string]string, error) {
	ret := _m.Called(storageType)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(kvstorage.Type) map[string]string); ok {
		r0 = rf(storageType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(kvstorage.Type) error); ok {
		r1 = rf(storageType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUnconfirmedTransactions provides a mock function with given fields:
func (_m *MockGatewayer) GetAllUnconfirmedTransactions() ([]visor.UnconfirmedTransaction, error) {
	ret := _m.Called()

	var r0 []visor.UnconfirmedTransaction
	if rf, ok := ret.Get(0).(func() []visor.UnconfirmedTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.UnconfirmedTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUnconfirmedTransactionsVerbose provides a mock function with given fields:
func (_m *MockGatewayer) GetAllUnconfirmedTransactionsVerbose() ([]visor.UnconfirmedTransaction, [][]visor.TransactionInput, error) {
	ret := _m.Called()

	var r0 []visor.UnconfirmedTransaction
	if rf, ok := ret.Get(0).(func() []visor.UnconfirmedTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.UnconfirmedTransaction)
		}
	}

	var r1 [][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func() [][]visor.TransactionInput); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBalanceOfAddresses provides a mock function with given fields: addrs
func (_m *MockGatewayer) GetBalanceOfAddresses(addrs []cipher.Address) ([]wallet.BalancePair, error) {
	ret := _m.Called(addrs)

	var r0 []wallet.BalancePair
	if rf, ok := ret.Get(0).(func([]cipher.Address) []wallet.BalancePair); ok {
		r0 = rf(addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]wallet.BalancePair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]cipher.Address) error); ok {
		r1 = rf(addrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockchainMetadata provides a mock function with given fields:
func (_m *MockGatewayer) GetBlockchainMetadata() (*visor.BlockchainMetadata, error) {
	ret := _m.Called()

	var r0 *visor.BlockchainMetadata
	if rf, ok := ret.Get(0).(func() *visor.BlockchainMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visor.BlockchainMetadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockchainProgress provides a mock function with given fields: headSeq
func (_m *MockGatewayer) GetBlockchainProgress(headSeq uint64) *daemon.BlockchainProgress {
	ret := _m.Called(headSeq)

	var r0 *daemon.BlockchainProgress
	if rf, ok := ret.Get(0).(func(uint64) *daemon.BlockchainProgress); ok {
		r0 = rf(headSeq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*daemon.BlockchainProgress)
		}
	}

	return r0
}

// GetBlocks provides a mock function with given fields: seqs
func (_m *MockGatewayer) GetBlocks(seqs []uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(seqs)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func([]uint64) []coin.SignedBlock); ok {
		r0 = rf(seqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint64) error); ok {
		r1 = rf(seqs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocksInRange provides a mock function with given fields: start, end
func (_m *MockGatewayer) GetBlocksInRange(start uint64, end uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(start, end)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64, uint64) []coin.SignedBlock); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlocksInRangeVerbose provides a mock function with given fields: start, end
func (_m *MockGatewayer) GetBlocksInRangeVerbose(start uint64, end uint64) ([]coin.SignedBlock, [][][]visor.TransactionInput, error) {
	ret := _m.Called(start, end)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64, uint64) []coin.SignedBlock); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 [][][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func(uint64, uint64) [][][]visor.TransactionInput); ok {
		r1 = rf(start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint64, uint64) error); ok {
		r2 = rf(start, end)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBlocksVerbose provides a mock function with given fields: seqs
func (_m *MockGatewayer) GetBlocksVerbose(seqs []uint64) ([]coin.SignedBlock, [][][]visor.TransactionInput, error) {
	ret := _m.Called(seqs)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func([]uint64) []coin.SignedBlock); ok {
		r0 = rf(seqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 [][][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func([]uint64) [][][]visor.TransactionInput); ok {
		r1 = rf(seqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]uint64) error); ok {
		r2 = rf(seqs)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetConnection provides a mock function with given fields: addr
func (_m *MockGatewayer) GetConnection(addr string) (*daemon.Connection, error) {
	ret := _m.Called(addr)

	var r0 *daemon.Connection
	if rf, ok := ret.Get(0).(func(string) *daemon.Connection); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*daemon.Connection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnections provides a mock function with given fields: f
func (_m *MockGatewayer) GetConnections(f func(daemon.Connection) bool) ([]daemon.Connection, error) {
	ret := _m.Called(f)

	var r0 []daemon.Connection
	if rf, ok := ret.Get(0).(func(func(daemon.Connection) bool) []daemon.Connection); ok {
		r0 = rf(f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]daemon.Connection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(func(daemon.Connection) bool) error); ok {
		r1 = rf(f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultConnections provides a mock function with given fields:
func (_m *MockGatewayer) GetDefaultConnections() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetExchgConnection provides a mock function with given fields:
func (_m *MockGatewayer) GetExchgConnection() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetLastBlocks provides a mock function with given fields: num
func (_m *MockGatewayer) GetLastBlocks(num uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(num)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64) []coin.SignedBlock); ok {
		r0 = rf(num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBlocksVerbose provides a mock function with given fields: num
func (_m *MockGatewayer) GetLastBlocksVerbose(num uint64) ([]coin.SignedBlock, [][][]visor.TransactionInput, error) {
	ret := _m.Called(num)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64) []coin.SignedBlock); ok {
		r0 = rf(num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 [][][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func(uint64) [][][]visor.TransactionInput); ok {
		r1 = rf(num)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRichlist provides a mock function with given fields: includeDistribution
func (_m *MockGatewayer) GetRichlist(includeDistribution bool) (visor.Richlist, error) {
	ret := _m.Called(includeDistribution)

	var r0 visor.Richlist
	if rf, ok := ret.Get(0).(func(bool) visor.Richlist); ok {
		r0 = rf(includeDistribution)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(visor.Richlist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(includeDistribution)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignedBlockByHash provides a mock function with given fields: hash
func (_m *MockGatewayer) GetSignedBlockByHash(hash cipher.SHA256) (*coin.SignedBlock, error) {
	ret := _m.Called(hash)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(cipher.SHA256) *coin.SignedBlock); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cipher.SHA256) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignedBlockByHashVerbose provides a mock function with given fields: hash
func (_m *MockGatewayer) GetSignedBlockByHashVerbose(hash cipher.SHA256) (*coin.SignedBlock, [][]visor.TransactionInput, error) {
	ret := _m.Called(hash)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(cipher.SHA256) *coin.SignedBlock); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 [][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func(cipher.SHA256) [][]visor.TransactionInput); ok {
		r1 = rf(hash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(cipher.SHA256) error); ok {
		r2 = rf(hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSignedBlockBySeq provides a mock function with given fields: seq
func (_m *MockGatewayer) GetSignedBlockBySeq(seq uint64) (*coin.SignedBlock, error) {
	ret := _m.Called(seq)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64) *coin.SignedBlock); ok {
		r0 = rf(seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(seq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignedBlockBySeqVerbose provides a mock function with given fields: seq
func (_m *MockGatewayer) GetSignedBlockBySeqVerbose(seq uint64) (*coin.SignedBlock, [][]visor.TransactionInput, error) {
	ret := _m.Called(seq)

	var r0 *coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64) *coin.SignedBlock); ok {
		r0 = rf(seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.SignedBlock)
		}
	}

	var r1 [][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func(uint64) [][]visor.TransactionInput); ok {
		r1 = rf(seq)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(seq)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSpentOutputsForAddresses provides a mock function with given fields: addr
func (_m *MockGatewayer) GetSpentOutputsForAddresses(addr []cipher.Address) ([][]historydb.UxOut, uint64, error) {
	ret := _m.Called(addr)

	var r0 [][]historydb.UxOut
	if rf, ok := ret.Get(0).(func([]cipher.Address) [][]historydb.UxOut); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]historydb.UxOut)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func([]cipher.Address) uint64); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]cipher.Address) error); ok {
		r2 = rf(addr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStorageValue provides a mock function with given fields: storageType, key
func (_m *MockGatewayer) GetStorageValue(storageType kvstorage.Type, key string) (string, error) {
	ret := _m.Called(storageType, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(kvstorage.Type, string) string); ok {
		r0 = rf(storageType, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(kvstorage.Type, string) error); ok {
		r1 = rf(storageType, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: txid
func (_m *MockGatewayer) GetTransaction(txid cipher.SHA256) (*visor.Transaction, error) {
	ret := _m.Called(txid)

	var r0 *visor.Transaction
	if rf, ok := ret.Get(0).(func(cipher.SHA256) *visor.Transaction); ok {
		r0 = rf(txid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visor.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cipher.SHA256) error); ok {
		r1 = rf(txid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionWithInputs provides a mock function with given fields: txid
func (_m *MockGatewayer) GetTransactionWithInputs(txid cipher.SHA256) (*visor.Transaction, []visor.TransactionInput, error) {
	ret := _m.Called(txid)

	var r0 *visor.Transaction
	if rf, ok := ret.Get(0).(func(cipher.SHA256) *visor.Transaction); ok {
		r0 = rf(txid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visor.Transaction)
		}
	}

	var r1 []visor.TransactionInput
	if rf, ok := ret.Get(1).(func(cipher.SHA256) []visor.TransactionInput); ok {
		r1 = rf(txid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(cipher.SHA256) error); ok {
		r2 = rf(txid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactions provides a mock function with given fields: flts, order, page
func (_m *MockGatewayer) GetTransactions(flts []visor.TxFilter, order visor.SortOrder, page *visor.PageIndex) ([]visor.Transaction, uint64, error) {
	ret := _m.Called(flts, order, page)

	var r0 []visor.Transaction
	if rf, ok := ret.Get(0).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) []visor.Transaction); ok {
		r0 = rf(flts, order, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.Transaction)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) uint64); ok {
		r1 = rf(flts, order, page)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) error); ok {
		r2 = rf(flts, order, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactionsWithInputs provides a mock function with given fields: flts, order, page
func (_m *MockGatewayer) GetTransactionsWithInputs(flts []visor.TxFilter, order visor.SortOrder, page *visor.PageIndex) ([]visor.Transaction, [][]visor.TransactionInput, uint64, error) {
	ret := _m.Called(flts, order, page)

	var r0 []visor.Transaction
	if rf, ok := ret.Get(0).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) []visor.Transaction); ok {
		r0 = rf(flts, order, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.Transaction)
		}
	}

	var r1 [][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) [][]visor.TransactionInput); ok {
		r1 = rf(flts, order, page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]visor.TransactionInput)
		}
	}

	var r2 uint64
	if rf, ok := ret.Get(2).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) uint64); ok {
		r2 = rf(flts, order, page)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func([]visor.TxFilter, visor.SortOrder, *visor.PageIndex) error); ok {
		r3 = rf(flts, order, page)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetTrustConnections provides a mock function with given fields:
func (_m *MockGatewayer) GetTrustConnections() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetUnspentOutputsSummary provides a mock function with given fields: filters
func (_m *MockGatewayer) GetUnspentOutputsSummary(filters []visor.OutputsFilter) (*visor.UnspentOutputsSummary, error) {
	ret := _m.Called(filters)

	var r0 *visor.UnspentOutputsSummary
	if rf, ok := ret.Get(0).(func([]visor.OutputsFilter) *visor.UnspentOutputsSummary); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visor.UnspentOutputsSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]visor.OutputsFilter) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUxOutByID provides a mock function with given fields: id
func (_m *MockGatewayer) GetUxOutByID(id cipher.SHA256) (*historydb.UxOut, uint64, error) {
	ret := _m.Called(id)

	var r0 *historydb.UxOut
	if rf, ok := ret.Get(0).(func(cipher.SHA256) *historydb.UxOut); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*historydb.UxOut)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(cipher.SHA256) uint64); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(cipher.SHA256) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWallet provides a mock function with given fields: wltID
func (_m *MockGatewayer) GetWallet(wltID string) (wallet.Wallet, error) {
	ret := _m.Called(wltID)

	var r0 wallet.Wallet
	if rf, ok := ret.Get(0).(func(string) wallet.Wallet); ok {
		r0 = rf(wltID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(wltID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletBalance provides a mock function with given fields: wltID
func (_m *MockGatewayer) GetWalletBalance(wltID string) (wallet.BalancePair, wallet.AddressBalances, error) {
	ret := _m.Called(wltID)

	var r0 wallet.BalancePair
	if rf, ok := ret.Get(0).(func(string) wallet.BalancePair); ok {
		r0 = rf(wltID)
	} else {
		r0 = ret.Get(0).(wallet.BalancePair)
	}

	var r1 wallet.AddressBalances
	if rf, ok := ret.Get(1).(func(string) wallet.AddressBalances); ok {
		r1 = rf(wltID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(wallet.AddressBalances)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(wltID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWalletSeed provides a mock function with given fields: wltID, password
func (_m *MockGatewayer) GetWalletSeed(wltID string, password []byte) (string, string, error) {
	ret := _m.Called(wltID, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, []byte) string); ok {
		r0 = rf(wltID, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, []byte) string); ok {
		r1 = rf(wltID, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte) error); ok {
		r2 = rf(wltID, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWalletUnconfirmedTransactions provides a mock function with given fields: wltID
func (_m *MockGatewayer) GetWalletUnconfirmedTransactions(wltID string) ([]visor.UnconfirmedTransaction, error) {
	ret := _m.Called(wltID)

	var r0 []visor.UnconfirmedTransaction
	if rf, ok := ret.Get(0).(func(string) []visor.UnconfirmedTransaction); ok {
		r0 = rf(wltID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.UnconfirmedTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(wltID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletUnconfirmedTransactionsVerbose provides a mock function with given fields: wltID
func (_m *MockGatewayer) GetWalletUnconfirmedTransactionsVerbose(wltID string) ([]visor.UnconfirmedTransaction, [][]visor.TransactionInput, error) {
	ret := _m.Called(wltID)

	var r0 []visor.UnconfirmedTransaction
	if rf, ok := ret.Get(0).(func(string) []visor.UnconfirmedTransaction); ok {
		r0 = rf(wltID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.UnconfirmedTransaction)
		}
	}

	var r1 [][]visor.TransactionInput
	if rf, ok := ret.Get(1).(func(string) [][]visor.TransactionInput); ok {
		r1 = rf(wltID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(wltID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWallets provides a mock function with given fields:
func (_m *MockGatewayer) GetWallets() (wallet.Wallets, error) {
	ret := _m.Called()

	var r0 wallet.Wallets
	if rf, ok := ret.Get(0).(func() wallet.Wallets); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Wallets)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadBkSeq provides a mock function with given fields:
func (_m *MockGatewayer) HeadBkSeq() (uint64, bool, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InjectBroadcastTransaction provides a mock function with given fields: txn
func (_m *MockGatewayer) InjectBroadcastTransaction(txn coin.Transaction) error {
	ret := _m.Called(txn)

	var r0 error
	if rf, ok := ret.Get(0).(func(coin.Transaction) error); ok {
		r0 = rf(txn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InjectTransaction provides a mock function with given fields: txn
func (_m *MockGatewayer) InjectTransaction(txn coin.Transaction) error {
	ret := _m.Called(txn)

	var r0 error
	if rf, ok := ret.Get(0).(func(coin.Transaction) error); ok {
		r0 = rf(txn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAddresses provides a mock function with given fields: wltID, password, n, options
func (_m *MockGatewayer) NewAddresses(wltID string, password []byte, n uint64, options ...wallet.Option) ([]cipher.Address, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, wltID, password, n)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []cipher.Address
	if rf, ok := ret.Get(0).(func(string, []byte, uint64, ...wallet.Option) []cipher.Address); ok {
		r0 = rf(wltID, password, n, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, uint64, ...wallet.Option) error); ok {
		r1 = rf(wltID, password, n, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecoverWallet provides a mock function with given fields: wltID, seed, seedPassphrase, password, options
func (_m *MockGatewayer) RecoverWallet(wltID string, seed string, seedPassphrase string, password []byte, options ...wallet.Option) (wallet.Wallet, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, wltID, seed, seedPassphrase, password)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 wallet.Wallet
	if rf, ok := ret.Get(0).(func(string, string, string, []byte, ...wallet.Option) wallet.Wallet); ok {
		r0 = rf(wltID, seed, seedPassphrase, password, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wallet.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, []byte, ...wallet.Option) error); ok {
		r1 = rf(wltID, seed, seedPassphrase, password, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveStorageValue provides a mock function with given fields: storageType, key
func (_m *MockGatewayer) RemoveStorageValue(storageType kvstorage.Type, key string) error {
	ret := _m.Called(storageType, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(kvstorage.Type, string) error); ok {
		r0 = rf(storageType, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResendUnconfirmedTxns provides a mock function with given fields:
func (_m *MockGatewayer) ResendUnconfirmedTxns() ([]cipher.SHA256, error) {
	ret := _m.Called()

	var r0 []cipher.SHA256
	if rf, ok := ret.Get(0).(func() []cipher.SHA256); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.SHA256)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanAddresses provides a mock function with given fields: wltID, password, n, tf
func (_m *MockGatewayer) ScanAddresses(wltID string, password []byte, n uint64, tf wallet.TransactionsFinder) ([]cipher.Address, error) {
	ret := _m.Called(wltID, password, n, tf)

	var r0 []cipher.Address
	if rf, ok := ret.Get(0).(func(string, []byte, uint64, wallet.TransactionsFinder) []cipher.Address); ok {
		r0 = rf(wltID, password, n, tf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, uint64, wallet.TransactionsFinder) error); ok {
		r1 = rf(wltID, password, n, tf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanWalletAddresses provides a mock function with given fields: wltID, password, num
func (_m *MockGatewayer) ScanWalletAddresses(wltID string, password []byte, num uint64) ([]cipher.Address, error) {
	ret := _m.Called(wltID, password, num)

	var r0 []cipher.Address
	if rf, ok := ret.Get(0).(func(string, []byte, uint64) []cipher.Address); ok {
		r0 = rf(wltID, password, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, uint64) error); ok {
		r1 = rf(wltID, password, num)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartedAt provides a mock function with given fields:
func (_m *MockGatewayer) StartedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// UnloadWallet provides a mock function with given fields: wltID
func (_m *MockGatewayer) UnloadWallet(wltID string) error {
	ret := _m.Called(wltID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(wltID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWalletLabel provides a mock function with given fields: wltID, label
func (_m *MockGatewayer) UpdateWalletLabel(wltID string, label string) error {
	ret := _m.Called(wltID, label)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(wltID, label)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyTxnVerbose provides a mock function with given fields: txn, signed
func (_m *MockGatewayer) VerifyTxnVerbose(txn *coin.Transaction, signed visor.TxnSignedFlag) ([]visor.TransactionInput, bool, error) {
	ret := _m.Called(txn, signed)

	var r0 []visor.TransactionInput
	if rf, ok := ret.Get(0).(func(*coin.Transaction, visor.TxnSignedFlag) []visor.TransactionInput); ok {
		r0 = rf(txn, signed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visor.TransactionInput)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*coin.Transaction, visor.TxnSignedFlag) bool); ok {
		r1 = rf(txn, signed)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*coin.Transaction, visor.TxnSignedFlag) error); ok {
		r2 = rf(txn, signed)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// VisorConfig provides a mock function with given fields:
func (_m *MockGatewayer) VisorConfig() visor.Config {
	ret := _m.Called()

	var r0 visor.Config
	if rf, ok := ret.Get(0).(func() visor.Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(visor.Config)
	}

	return r0
}

// WalletCreateTransaction provides a mock function with given fields: wltID, p, wp
func (_m *MockGatewayer) WalletCreateTransaction(wltID string, p transaction.Params, wp visor.CreateTransactionParams) (*coin.Transaction, []visor.TransactionInput, error) {
	ret := _m.Called(wltID, p, wp)

	var r0 *coin.Transaction
	if rf, ok := ret.Get(0).(func(string, transaction.Params, visor.CreateTransactionParams) *coin.Transaction); ok {
		r0 = rf(wltID, p, wp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.Transaction)
		}
	}

	var r1 []visor.TransactionInput
	if rf, ok := ret.Get(1).(func(string, transaction.Params, visor.CreateTransactionParams) []visor.TransactionInput); ok {
		r1 = rf(wltID, p, wp)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, transaction.Params, visor.CreateTransactionParams) error); ok {
		r2 = rf(wltID, p, wp)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// WalletCreateTransactionSigned provides a mock function with given fields: wltID, password, p, wp
func (_m *MockGatewayer) WalletCreateTransactionSigned(wltID string, password []byte, p transaction.Params, wp visor.CreateTransactionParams) (*coin.Transaction, []visor.TransactionInput, error) {
	ret := _m.Called(wltID, password, p, wp)

	var r0 *coin.Transaction
	if rf, ok := ret.Get(0).(func(string, []byte, transaction.Params, visor.CreateTransactionParams) *coin.Transaction); ok {
		r0 = rf(wltID, password, p, wp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.Transaction)
		}
	}

	var r1 []visor.TransactionInput
	if rf, ok := ret.Get(1).(func(string, []byte, transaction.Params, visor.CreateTransactionParams) []visor.TransactionInput); ok {
		r1 = rf(wltID, password, p, wp)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte, transaction.Params, visor.CreateTransactionParams) error); ok {
		r2 = rf(wltID, password, p, wp)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// WalletDir provides a mock function with given fields:
func (_m *MockGatewayer) WalletDir() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WalletSignTransaction provides a mock function with given fields: wltID, password, txn, signIndexes
func (_m *MockGatewayer) WalletSignTransaction(wltID string, password []byte, txn *coin.Transaction, signIndexes []int) (*coin.Transaction, []visor.TransactionInput, error) {
	ret := _m.Called(wltID, password, txn, signIndexes)

	var r0 *coin.Transaction
	if rf, ok := ret.Get(0).(func(string, []byte, *coin.Transaction, []int) *coin.Transaction); ok {
		r0 = rf(wltID, password, txn, signIndexes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coin.Transaction)
		}
	}

	var r1 []visor.TransactionInput
	if rf, ok := ret.Get(1).(func(string, []byte, *coin.Transaction, []int) []visor.TransactionInput); ok {
		r1 = rf(wltID, password, txn, signIndexes)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]visor.TransactionInput)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte, *coin.Transaction, []int) error); ok {
		r2 = rf(wltID, password, txn, signIndexes)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
