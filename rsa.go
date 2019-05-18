package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type RsaPublicKey struct {
	Exp, Mod *big.Int
}

// Perform RSA encryption as defined in RSA.js (v1) developed by David Shapiro.
// The JavaScript is saved to ./relatedres/RSA.js
// Official website: http://www.ohdave.com/rsa/
func RsaEncrypt(pubKey *RsaPublicKey, msg string, radix int) string {
	if pubKey == nil {
		panic(errors.New("ucasnauth: RSA public key is nil"))
	}
	if pubKey.Exp == nil || pubKey.Mod == nil {
		panic(errors.New("ucasnauth: RSA public key is invalid"))
	}
	if radix < 2 || radix > int(big.MaxBase) {
		panic(fmt.Errorf("ucasnauth: radix must be between 2 and %d, inclusive.",
			big.MaxBase))
	}
	msgBytes := []byte(msg)
	cs := pubKey.ChunkSize()
	if cs <= 0 {
		panic(fmt.Errorf("ucasnauth: chunk size (%d) is non-positive", cs))
	}
	data := make([]byte, len(msgBytes)+cs-1-(len(msgBytes)+cs-1)%cs)
	dataLen := len(data)
	// Reverse bytes of msg because big.Int is big-endian but David's BigInt is little-endian.
	for i := range msgBytes {
		data[dataLen-i-1] = msgBytes[i]
	}
	chunk := new(big.Int)
	var builder strings.Builder
	dataLenSubCs := dataLen - cs
	for i := 0; i < dataLen; i += cs {
		chunk.SetBytes(data[i:i+cs]).Exp(chunk, pubKey.Exp, pubKey.Mod)
		s := chunk.Text(radix)
		if i < dataLenSubCs {
			builder.Grow(len(s) + 1)
		}
		builder.WriteString(strings.ToLower(s))
		if i < dataLenSubCs {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func NewRsaPublicKey(exp int, modStr string, modBase int) *RsaPublicKey {
	mod := new(big.Int)
	_, ok := mod.SetString(modStr, modBase)
	if !ok {
		panic(errors.New("ucasnauth: invalid modStr of modBase"))
	}
	return &RsaPublicKey{
		Exp: big.NewInt(int64(exp)),
		Mod: mod,
	}
}

func (rpk *RsaPublicKey) ChunkSize() int {
	if rpk == nil || rpk.Mod == nil {
		return 0
	}
	cs := rpk.Mod.BitLen() - 1
	if cs < 0 {
		cs = 0
	}
	return cs
}
