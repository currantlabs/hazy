package hazy

import (
	"crypto/rand"
	"math/big"
	"testing"
	"time"
)

func TestEncodeDecode(t *testing.T) {

	b := make([]byte, 8)
	now := time.Now()
	for i := 0; i < 10000; i++ {
		rand.Read(b)
		u := new(big.Int).SetBytes(b).Uint64()
		s := encode(u)
		println(string(s))
		du, err := decode(s)
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
