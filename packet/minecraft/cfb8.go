package minecraft

import "crypto/cipher"

type cfb8 struct {
	block			cipher.Block
	iv				[]byte
	tmp				[]byte
	decrypt			bool
}

func NewCFB8Decrypter(block cipher.Block, iv []byte) (stream cipher.Stream) {
	return newCFB8(block, iv, true)
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) (stream cipher.Stream){
	return newCFB8(block, iv, false)
}

func newCFB8(block cipher.Block, iv []byte, decrypt bool) (stream cipher.Stream) {
	bytes := make([]byte, len(iv))
	copy(bytes, iv)
	return &cfb8{
		block:		block,
		iv:			bytes,
		tmp:		make([]byte, block.BlockSize()),
		decrypt:	decrypt,
	}
}

func (this *cfb8) XORKeyStream(dst, src []byte) {
	var val byte
	for i := 0; i < len(src); i++ {
		val = src[i]
		copy(this.tmp, this.iv)
		this.block.Encrypt(this.iv, this.iv)
		val = val ^ this.iv[0]
		copy(this.iv, this.tmp[1:]);
		if this.decrypt {
			this.iv[15] = src[i]
		} else {
			this.iv[15] = val
		}
		dst[i] = val
	}
}
