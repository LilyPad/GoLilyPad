package proxy

import "crypto/rand"
import "crypto/sha1"
import "encoding/hex"
import "io"
import "github.com/LilyPad/GoLilyPad/packet/minecraft"

func MinecraftVersion() string {
	return minecraft.STRING_VERSION
}

func RandomBytes(size int) (bytes []byte, err error) {
	bytes = make([]byte, size)
	_, err = io.ReadFull(rand.Reader, bytes)
	return
}

func GenSalt() (str string, err error) {
	salt := make([]byte, 10)
	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return
	}
	str = hex.EncodeToString(salt)
	return
}

func Sha1Hex(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum(nil))
}