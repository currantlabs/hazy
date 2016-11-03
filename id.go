package hazy

import (
	"math/big"
	"sync"
)

type ID struct {
	Clear uint64
	Hazy  uint64
}

const IDLength = 13

var Prime *big.Int
var Coprime *big.Int
var Pepper uint64
var uint64Max = new(big.Int).SetUint64(18446744073709551615)

var pool = sync.Pool{
	New: func() interface{} {
		return new(big.Int)
	},
}

func Initialize(prime uint64, coprime uint64, pepper uint64) error {
	Prime = new(big.Int).SetUint64(prime)
	if !Prime.ProbablyPrime(40) {
		return ErrInvalidPrime
	}
	Coprime = new(big.Int).SetUint64(coprime)
	Pepper = pepper
	return nil
}

func (id ID) IsZero() bool {
	return id.Clear == 0
}

func (id ID) Equal(other ID) bool {
	return id.Clear == other.Clear
}

func (id ID) String() string {
	return string(Base32Encode(id.Hazy))
}

func Obscure(id uint64) ID {
	return ID{
		Clear: id,
		Hazy:  obscure(id),
	}
}

func Reveal(id uint64) ID {
	return ID{
		Clear: reveal(id),
		Hazy:  id,
	}
}

func obscure(id uint64) uint64 {
	i := pool.Get().(*big.Int)
	i.SetUint64(id)
	i.Mul(i, Prime)
	i.And(i, uint64Max)
	id = i.Uint64() ^ Pepper
	pool.Put(i)
	return id
}

func reveal(id uint64) uint64 {
	i := pool.Get().(*big.Int)
	i.SetUint64(id ^ Pepper)
	i.Mul(i, Coprime)
	i.And(i, uint64Max)
	id = i.Uint64()
	pool.Put(i)
	return id
}

func ObscureWithPrime(id uint64, prime uint64, pepper uint64) (uint64, error) {
	i := pool.Get().(*big.Int)
	i.SetUint64(id)
	p := pool.Get().(*big.Int)
	p.SetUint64(prime)
	if !p.ProbablyPrime(40) {
		return 0, ErrInvalidPrime
	}
	i.Mul(i, p)
	i.And(i, uint64Max)
	id = i.Uint64() ^ pepper
	pool.Put(i)
	pool.Put(p)
	return id, nil
}

func RevealWithCoprime(id uint64, coprime uint64, pepper uint64) uint64 {
	i := pool.Get().(*big.Int)
	i.SetUint64(id ^ pepper)
	cp := pool.Get().(*big.Int)
	cp.SetUint64(coprime)
	i.Mul(i, cp)
	i.And(i, uint64Max)
	id = i.Uint64()
	pool.Put(i)
	return id
}
