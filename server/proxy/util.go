package proxy

import (
	cryptoRand "crypto/rand"
	"math/rand"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
)

func MinecraftVersion() string {
	return minecraft.STRING_VERSION
}

func RandomBytes(size int) (bytes []byte, err error) {
	bytes = make([]byte, size)
	_, err = io.ReadFull(cryptoRand.Reader, bytes)
	return
}

func RandomInt(max int) (val int) {
	val = rand.Intn(max)
	return
}

func FormatUUID(uuid string) (val string) {
	if len(uuid) == 32 {
		val = uuid[:8] + "-" + uuid[8:12] + "-" + uuid[12:16] + "-" + uuid[16:20] + "-" + uuid[20:]
	}
	return
}

func GenNameUUID(name string) (val string) {
	md5 := md5.New()
	md5.Write([]byte(name))
	md5Sum := md5.Sum(nil)
	md5Sum[6] &= 0x0F
	md5Sum[6] |= 0x30
	md5Sum[8] &= 0x3F
	md5Sum[8] |= 0x80
	val = hex.EncodeToString(md5Sum)
	return
}

func GenSalt() (str string, err error) {
	salt := make([]byte, 10)
	_, err = io.ReadFull(cryptoRand.Reader, salt)
	if err != nil {
		return
	}
	str = hex.EncodeToString(salt)
	return
}

func Sha1Hex(str string) (val string) {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	val = hex.EncodeToString(sha1.Sum(nil))
	return
}
