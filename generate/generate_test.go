package generate

import (
	"testing"

	"github.com/currantlabs/hazy"
	"github.com/pkg/errors"
)

func TestGenerate(t *testing.T) {
	prime, coprime, pepper, err := New()
	if err != nil {
		t.Error(err)
	}
	id := randomUint64()
	obs, err := hazy.ObscureWithPrime(id, prime, pepper)
	if err != nil {
		t.Error(err)
	}
	rev := hazy.RevealWithCoprime(obs, coprime, pepper)
	if id != rev {
		t.Error(errors.New("mismatch reveal obscure"))
	}
}
