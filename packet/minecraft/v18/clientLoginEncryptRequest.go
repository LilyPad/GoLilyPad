package v18

import (
	"errors"
	"fmt"
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
	publicKeyLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if publicKeyLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Public key length is below zero: %d", publicKeyLength))
		return
	}
	if publicKeyLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Public key length is above maximum: %d", publicKeyLength))
		return
	}
	packetClientLoginEncryptRequest.PublicKey = make([]byte, publicKeyLength)
	_, err = reader.Read(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	verifyTokenLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if verifyTokenLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Verify token length is below zero: %d", verifyTokenLength))
		return
	}
	if verifyTokenLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Verify token length is above maximum: %d", verifyTokenLength))
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
	err = packet.WriteVarInt(writer, len(packetClientLoginEncryptRequest.PublicKey))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(packetClientLoginEncryptRequest.VerifyToken))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.VerifyToken)
	return
}
