
#include <stdlib.h>
#include <time.h>

#include <criterion/criterion.h>
#include "libskycoin.h"
#include "skyerrors.h"

#define ALPHANUM "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
#define ALPHANUM_LEN 62

void randBytes(unsigned char* bytes, size_t n) {
  size_t i;
  unsigned char *ptr;
  for (i = 0, ptr = bytes; i < n; ++i, ++ptr) {
    *ptr = ALPHANUM[rand() % ALPHANUM_LEN];
  } 
}

void setup(void) {
  srand ((unsigned int) time (NULL));
}

Test(asserts, TestNewPubKey) {
  // buffer big enough to hold all kind of data needed by test cases
  unsigned char buff[101];
  GoSlice slice;
  PubKey pk;

  slice.data = buff;
  slice.cap = 101;

  randBytes(buff, 31);
  slice.len = 31;

  unsigned int errcode = SKY_cipher_NewPubKey(slice, &pk);
  cr_assert(errcode == SKY_ERROR, "33 random bytes");
}

