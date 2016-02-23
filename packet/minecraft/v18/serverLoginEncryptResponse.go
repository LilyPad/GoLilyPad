package v18

import (
	"errors"
	"fmt"
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
	sharedSecretLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if sharedSecretLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Shared secret length is below zero: %d", sharedSecretLength))
		return
	}
	if sharedSecretLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Shared secret length is above maximum: %d", sharedSecretLength))
		return
	}
	packetServerLoginEncryptResponse.SharedSecret = make([]byte, sharedSecretLength)
	_, err = reader.Read(packetServerLoginEncryptResponse.SharedSecret)
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
	err = packet.WriteVarInt(writer, len(packetServerLoginEncryptResponse.SharedSecret))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(packetServerLoginEncryptResponse.VerifyToken))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.VerifyToken)
	return
}
