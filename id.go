package hazy

import (
	"math/big"
)

type ID struct {
	Clear uint64
	Hazy  uint64
}

var Prime *big.Int
var Coprime *big.Int
var Pepper uint64
var uint64Max = new(big.Int).SetUint64(18446744073709551615)

func Initialize(prime uint64, coprime uint64, pepper uint64) error {
	Prime = new(big.Int).SetUint64(prime)
	if !Prime.ProbablyPrime(40) {
		return ErrInvalidPrime
	}
	Coprime = new(big.Int).SetUint64(coprime)
	if !Coprime.ProbablyPrime(40) {
		return ErrInvalidCoprime
	}
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
	return string(encode(id.Hazy))
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
	var i big.Int
	i.SetUint64(id)
	i.Mul(&i, Prime)
	i.And(&i, uint64Max)
	id = i.Uint64() ^ Pepper
	return id
}

func reveal(id uint64) uint64 {
	var i big.Int
	i.SetUint64(id ^ Pepper)
	i.Mul(&i, Coprime)
	i.And(&i, uint64Max)
	id = i.Uint64()
	return id
}
