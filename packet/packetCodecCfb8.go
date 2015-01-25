package packet

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

type PacketCodecCfb8 struct {
	codec PacketCodec
	writer *cipher.StreamWriter
	reader *cipher.StreamReader
}

func NewPacketCodecCfb8(sharedSecret []byte) (this *PacketCodecCfb8, err error) {
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return
	}
	this = new(PacketCodecCfb8)
	this.writer = new(cipher.StreamWriter)
	this.writer.S = NewCFB8Encrypt(block, sharedSecret)
	this.reader = new(cipher.StreamReader)
	this.reader.S = NewCFB8Decrypt(block, sharedSecret)
	return
}

func (this *PacketCodecCfb8) Decode(reader io.Reader) (packet Packet, err error) {
	this.reader.R = reader
	packet, err = this.codec.Decode(this.reader)
	return
}

func (this *PacketCodecCfb8) Encode(writer io.Writer, packet Packet) (err error) {
	this.writer.W = writer
	err = this.codec.Encode(this.writer, packet)
	return
}

func (this *PacketCodecCfb8) SetCodec(codec PacketCodec) {
	this.codec = codec
}
