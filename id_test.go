package hazy

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"
)

func TestObscureReveal(t *testing.T) {
	Initialize(0xf1b62104017f2969, 0xc5a3e2ec4db3f6d9, 0xfe5fb8ad8031c686)
	var i uint64
	for i = 0; i < 10000; i++ {
		di := Reveal(Obscure(i).Hazy)
		if di.Clear != i {
			t.Errorf("obscure reveal mismatch %v %v", i, di)
		}
	}
}

func BenchmarkObscure(b *testing.B) {
	Initialize(0xf1b62104017f2969, 0xc5a3e2ec4db3f6d9, 0xfe5fb8ad8031c686)
	for i := 1; i < b.N; i++ {
		Obscure(uint64(i))
	}
}

func BenchmarkObscureReveal(b *testing.B) {
	Initialize(0xf1b62104017f2969, 0xc5a3e2ec4db3f6d9, 0xfe5fb8ad8031c686)
	for i := 1; i < b.N; i++ {
		Reveal(Obscure(uint64(i)).Hazy)
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	ub := make([]byte, 8)
	rand.Read(ub)
	u := new(big.Int).SetBytes(ub).Uint64()
	for i := 1; i < b.N; i++ {
		du, err := Base32Decode(Base32Encode(u))
		if err != nil {
			b.Error("failed decoding %v %v", u, err)
			return
		}
		if u != du {
			b.Error("decode mismatch %v %v", u, du)
			return
		}
	}
}

func BenchmarkEncodeDecodeHex(b *testing.B) {
	ub := make([]byte, 8)
	rand.Read(ub)
	for i := 1; i < b.N; i++ {
		hex.DecodeString(hex.EncodeToString(ub))
	}
}
