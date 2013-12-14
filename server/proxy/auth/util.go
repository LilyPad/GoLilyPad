package auth

import "crypto/sha1"
import "encoding/hex"
import "strings"

func MojangSha1Hex(arrBytes ...[]byte) string {
	sha1 := sha1.New()
	for _, bytes := range arrBytes {
		sha1.Write(bytes)
	}
	hash := sha1.Sum(nil)
	negative := (hash[0] & 0x80) == 0x80
	if negative {
		twosCompliment(hash)
	}
	hexString := hex.EncodeToString(hash)
	hexString = strings.TrimLeft(hexString, "0")
	if negative {
		hexString = "-" + hexString
	}
	return hexString
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
