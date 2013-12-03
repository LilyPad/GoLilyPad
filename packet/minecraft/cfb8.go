package minecraft

import "crypto/cipher"
import "errors"

type cfb8 struct {
    block               cipher.Block
    iv, initialIv, tmp 	[]byte
    decrypt             bool
}

func NewCFB8Decrypter(block cipher.Block, iv []byte) (stream cipher.Stream, err error) {
    return newCFB8(block, iv, true)
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) (stream cipher.Stream, err error) {
    return newCFB8(block, iv, false)
}

func newCFB8(block cipher.Block, iv []byte, decrypt bool) (stream cipher.Stream, err error) {
    if len(iv) != 16 {
    	err = errors.New("IV not valid")
    	return
    }
    bytes := make([]byte, 256)
    copy(bytes, iv)
    return &cfb8{
        block:     block,
        iv:        bytes[:16],
        initialIv: bytes,
        tmp:       make([]byte, 16),
        decrypt:   decrypt,
    }, nil
}

func (this *cfb8) XORKeyStream(dst, src []byte) {
	var val byte
    for i := 0; i < len(src); i++ {
        this.block.Encrypt(this.tmp, this.iv)
        val = src[i] ^ this.tmp[0]
        if cap(this.iv) >= 17 {
                this.iv = this.iv[1:17]
        } else {
                copy(this.initialIv, this.iv[1:])
                this.iv = this.initialIv[:16]
        }
        if this.decrypt {
                this.iv[15] = src[i]
        } else {
                this.iv[15] = val
        }
        dst[i] = val
    }
}