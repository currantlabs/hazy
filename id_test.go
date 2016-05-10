package hazy

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"
	"time"
)

func TestObscureReveal(t *testing.T) {

	b := make([]byte, 8)
	now := time.Now()
	for i := 0; i < 10000; i++ {
		rand.Read(b)
		u := new(big.Int).SetBytes(b).Uint64()
		s := Base32Encode(u)
		//println(string(s))
		du, err := Base32Decode(s)
		if err != nil {
			t.Error("failed decoding %v %v", u, err)
			return
		}
		if u != du {
			t.Error("decode mismatch", u, du, string(s))
			return
		}
	}
	println(time.Duration(time.Now().Sub(now).Nanoseconds() / 10000).String())
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
		Base32Decode(Base32Encode(u))
	}
}

func BenchmarkEncodeDecodeHex(b *testing.B) {
	ub := make([]byte, 8)
	rand.Read(ub)
	for i := 1; i < b.N; i++ {
		hex.DecodeString(hex.EncodeToString(ub))
	}
}
