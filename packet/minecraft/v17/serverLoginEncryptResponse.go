package v17

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecServerLoginEncryptResponse struct {
	IdMap *minecraft.IdMap
}

func (this *CodecServerLoginEncryptResponse) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerLoginEncryptResponse := new(minecraft.PacketServerLoginEncryptResponse)
	packetServerLoginEncryptResponse.IdFrom(this.IdMap)
	sharedSecretLength, err := packet.ReadUint16(reader)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.SharedSecret = make([]byte, sharedSecretLength)
	_, err = reader.Read(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	verifyTokenLength, err := packet.ReadUint16(reader)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.VerifyToken = make([]byte, verifyTokenLength)
	_, err = reader.Read(packetServerLoginEncryptResponse.VerifyToken)
	if err != nil {
		return
	}
	decode = packetServerLoginEncryptResponse
	return
}

func (this *CodecServerLoginEncryptResponse) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerLoginEncryptResponse := encode.(*minecraft.PacketServerLoginEncryptResponse)
	err = packet.WriteUint16(writer, uint16(len(packetServerLoginEncryptResponse.SharedSecret)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, uint16(len(packetServerLoginEncryptResponse.VerifyToken)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.VerifyToken)
	return
}
