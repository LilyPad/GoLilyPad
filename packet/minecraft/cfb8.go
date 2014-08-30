package minecraft

import (
	"crypto/cipher"
)

type cfb8 struct {
	Block cipher.Block
	Iv []byte
	Tmp []byte
	Decrypt bool
}
func newCFB8(block cipher.Block, iv []byte, decrypt bool) (stream cipher.Stream) {
	cfb8 := new(cfb8)
	cfb8.Block = block
	cfb8.Iv = make([]byte, len(iv))
	cfb8.Tmp = make([]byte, block.BlockSize())
	cfb8.Decrypt = decrypt
	copy(cfb8.Iv, iv)
	stream = cfb8
	return
}

func NewCFB8Encrypt(block cipher.Block, iv []byte) (stream cipher.Stream) {
	stream = newCFB8(block, iv, false)
	return
}

func NewCFB8Decrypt(block cipher.Block, iv []byte) (stream cipher.Stream) {
	stream = newCFB8(block, iv, true)
	return
}

func (this *cfb8) XORKeyStream(dst, src []byte) {
	var val byte
	for i := 0; i < len(src); i++ {
		val = src[i]
		copy(this.Tmp, this.Iv)
		this.Block.Encrypt(this.Iv, this.Iv)
		val = val ^ this.Iv[0]
		copy(this.Iv, this.Tmp[1:]);
		if this.Decrypt {
			this.Iv[15] = src[i]
		} else {
			this.Iv[15] = val
		}
		dst[i] = val
	}
}
