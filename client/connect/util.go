package connect

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Hex(str string) (val string) {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	val = hex.EncodeToString(sha1.Sum(nil))
	return
}

func PasswordAndSaltHash(password string, passwordSalt string) (val string) {
	val = Sha1Hex(Sha1Hex(passwordSalt) + Sha1Hex(password))
	return
}
