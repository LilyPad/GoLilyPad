package minecraft

import (
	"errors"
	"fmt"
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientLoginEncryptRequest struct {
	ServerId string
	PublicKey []byte
	VerifyToken []byte
}

func NewPacketClientLoginEncryptRequest(serverId string, publicKey []byte, verifyToken []byte) (this *PacketClientLoginEncryptRequest) {
	this = new(PacketClientLoginEncryptRequest)
	this.ServerId = serverId
	this.PublicKey = publicKey
	this.VerifyToken = verifyToken
	return
}

func (this *PacketClientLoginEncryptRequest) Id() int {
	return PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST
}

type packetClientLoginEncryptRequestCodec struct {

}

func (this *packetClientLoginEncryptRequestCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginEncryptRequest := new(PacketClientLoginEncryptRequest)
	packetClientLoginEncryptRequest.ServerId, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	publicKeyLength, err := packet.ReadVarInt(reader, util)
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
	verifyTokenLength, err := packet.ReadVarInt(reader, util)
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

func (this *packetClientLoginEncryptRequestCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginEncryptRequest := encode.(*PacketClientLoginEncryptRequest)
	err = packet.WriteString(writer, util, packetClientLoginEncryptRequest.ServerId)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, util, len(packetClientLoginEncryptRequest.PublicKey))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, util, len(packetClientLoginEncryptRequest.VerifyToken))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.VerifyToken)
	return
}
