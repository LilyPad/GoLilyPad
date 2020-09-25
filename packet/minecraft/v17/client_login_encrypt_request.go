package v17

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientLoginEncryptRequest struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientLoginEncryptRequest) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientLoginEncryptRequest := new(minecraft.PacketClientLoginEncryptRequest)
	packetClientLoginEncryptRequest.IdFrom(this.IdMap)
	packetClientLoginEncryptRequest.ServerId, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	publicKeyLength, err := packet.ReadUint16(reader)
	if err != nil {
		return
	}
	packetClientLoginEncryptRequest.PublicKey = make([]byte, publicKeyLength)
	_, err = reader.Read(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	verifyTokenLength, err := packet.ReadUint16(reader)
	if err != nil {
		return
	}
	packetClientLoginEncryptRequest.VerifyToken = make([]byte, verifyTokenLength)
	_, err = reader.Read(packetClientLoginEncryptRequest.VerifyToken)
	if err != nil {
		return
	}
	decode = packetClientLoginEncryptRequest
	return
}

func (this *CodecClientLoginEncryptRequest) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientLoginEncryptRequest := encode.(*minecraft.PacketClientLoginEncryptRequest)
	err = packet.WriteString(writer, packetClientLoginEncryptRequest.ServerId)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, uint16(len(packetClientLoginEncryptRequest.PublicKey)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, uint16(len(packetClientLoginEncryptRequest.VerifyToken)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.VerifyToken)
	return
}
