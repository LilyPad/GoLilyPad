package connect

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"io"
)

func RandomInt() (n int32) {
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return
}

func GenSalt() (str string, err error) {
	salt := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return
	}
	str = hex.EncodeToString(salt)
	return
}

func GenUUID() (str string, err error) {
	uuid := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, uuid)
	if err != nil {
		return
	}
	uuid[8] = 0x80
	uuid[4] = 0x40
	str = hex.EncodeToString(uuid)
	return
}

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
