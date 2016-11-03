package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/currantlabs/hazy"
	"github.com/currantlabs/hazy/generate"
)

type result struct {
	Prime      string
	Coprime    string
	Pepper     string
	Zero       string
	PrimeHex   string
	CoprimeHex string
	PepperHex  string
}

var resultTemp = template.Must(template.New("result").Parse(
	`Hazy parameter generation complete!

Prime: {{.Prime}}
Coprime: {{.Coprime}}
Pepper: {{.Pepper}}
Hazy Zero: {{.Zero}}

Paste the following into your go program before using hazy.ID:

import (
	"math/big"

	"github.com/currantlabs/hazy"
)

hazy.Initialize(0x{{.PrimeHex}}, 0x{{.CoprimeHex}}, 0x{{.PepperHex}})

`))

func main() {
	prime, coprime, pepper, err := generate.New()
	if err != nil {
		panic(err)
	}
	hazyZero, _ := hazy.ObscureWithPrime(0, prime, pepper)
	resultTemp.Execute(os.Stdout, &result{
		Prime:      fmt.Sprintf("%v", prime),
		Coprime:    fmt.Sprintf("%v", coprime),
		Pepper:     fmt.Sprintf("%v", pepper),
		Zero:       fmt.Sprintf("%v", hazyZero),
		PrimeHex:   fmt.Sprintf("%x", prime),
		CoprimeHex: fmt.Sprintf("%x", coprime),
		PepperHex:  fmt.Sprintf("%x", pepper),
	})
}
