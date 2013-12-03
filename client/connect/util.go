package connect

import "crypto/sha1"
import "encoding/hex"

func Sha1Hex(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum(nil))
}

func PasswordAndSaltHash(password string, passwordSalt string) string {
	return Sha1Hex(Sha1Hex(passwordSalt) + Sha1Hex(password))
}