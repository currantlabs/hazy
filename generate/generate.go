package generate

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/currantlabs/hazy"
	"github.com/cznic/mathutil"
	"github.com/pkg/errors"
)

func New() (uint64, uint64, uint64, error) {
	prime, err := generatePrime()
	if err != nil {
		return 0, 0, 0, err
	}
	coprime := getCoprime(prime)
	pepper := randomUint64()
	var testVals = []uint64{1, prime, coprime, 18446744073709551615}
	for _, val := range testVals {
		err = test(val, prime, coprime, pepper)
		if err != nil {
			return 0, 0, 0, err
		}
	}
	return prime, coprime, pepper, nil
}

func generatePrime() (uint64, error) {
	for i := 0; i < 100000; i++ {
		prime := randomUint64()
		if mathutil.IsPrimeUint64(prime) {
			return prime, nil
		}
	}
	return 0, errors.New("timed out generating prime")
}

func randomUint64() uint64 {
	b := make([]byte, 8)
	rand.Read(b)
	return new(big.Int).SetBytes(b).Uint64()
}

func getCoprime(prime uint64) uint64 {
	coprime := new(big.Int).SetUint64(prime)
	mod := new(big.Int).SetBytes([]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}) // 2^64
	return coprime.ModInverse(coprime, mod).Uint64()
}

func test(val uint64, prime uint64, coprime uint64, pepper uint64) error {
	obs, err := hazy.ObscureWithPrime(val, prime, pepper)
	if err != nil {
		return err
	}
	clear := hazy.RevealWithCoprime(obs, coprime, pepper)
	if clear != val {
		return fmt.Errorf("mismatch! %v %v", val, clear)
	}
	return nil
}
