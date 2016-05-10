package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"text/template"

	"github.com/currantlabs/hazy"
	"github.com/cznic/mathutil"
)

type result struct {
	Prime      string
	Coprime    string
	Pepper     string
	PrimeHex   string
	CoprimeHex string
	PepperHex  string
}

var resultTemp = template.Must(template.New("result").Parse(
	`Hazy parameter generation complete!

Prime: {{.Prime}}
Coprime: {{.Coprime}}
Pepper: {{.Pepper}}

Paste the following into your go program before using hazy.ID:

import (
	"math/big"

	"github.com/currantlabs/hazy"
)

hazy.Initialize(0x{{.PrimeHex}}, 0x{{.CoprimeHex}}, 0x{{.PepperHex}})

`))

func main() {
	prime := generatePrime()
	println(prime)
	coprime := getCoprime(prime)
	println(coprime)
	pepper := randomUint64()
	test(1, prime, coprime, pepper)
	test(prime, prime, coprime, pepper)
	test(coprime, prime, coprime, pepper)
	test(18446744073709551615, prime, coprime, pepper)
	resultTemp.Execute(os.Stdout, &result{
		Prime:      fmt.Sprintf("%v", prime),
		Coprime:    fmt.Sprintf("%v", coprime),
		Pepper:     fmt.Sprintf("%v", pepper),
		PrimeHex:   fmt.Sprintf("%x", prime),
		CoprimeHex: fmt.Sprintf("%x", coprime),
		PepperHex:  fmt.Sprintf("%x", pepper),
	})
}

func generatePrime() uint64 {
	for {
		prime := randomUint64()
		if mathutil.IsPrimeUint64(prime) {
			return prime
		}
	}
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

func test(val uint64, prime uint64, coprime uint64, pepper uint64) {
	hazy.Initialize(prime, coprime, pepper)
	val2 := hazy.Reveal(hazy.Obscure(val).Hazy).Clear
	if val != val2 {
		panic(fmt.Errorf("mismatch! %v %v", val, val2))
	}
}
