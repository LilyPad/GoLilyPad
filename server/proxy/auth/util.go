package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func MojangSha1Hex(arrBytes ...[]byte) (val string) {
	sha1 := sha1.New()
	for _, bytes := range arrBytes {
		sha1.Write(bytes)
	}
	hash := sha1.Sum(nil)
	negative := (hash[0] & 0x80) == 0x80
	if negative {
		twosCompliment(hash)
	}
	val = hex.EncodeToString(hash)
	val = strings.TrimLeft(val, "0")
	if negative {
		val = "-" + val
	}
	return
}

func twosCompliment(p []byte) {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = ^p[i]
		if carry {
			carry = p[i] == 0xFF
			p[i]++
		}
	}
}
