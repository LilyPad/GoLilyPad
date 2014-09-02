package minecraft

import (
	"errors"
	"fmt"
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerLoginEncryptResponse struct {
	SharedSecret []byte
	VerifyToken []byte
}

func NewPacketServerLoginEncryptResponse(sharedSecret []byte, verifyToken []byte) (this *PacketServerLoginEncryptResponse) {
	this = new(PacketServerLoginEncryptResponse)
	this.SharedSecret = sharedSecret
	this.VerifyToken = verifyToken
	return
}

func (this *PacketServerLoginEncryptResponse) Id() int {
	return PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE
}

type packetServerLoginEncryptResponseCodec struct {

}

func (this *packetServerLoginEncryptResponseCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerLoginEncryptResponse := new(PacketServerLoginEncryptResponse)
	sharedSecretLength, err := packet.ReadVarInt(reader, util)
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
	packetServerLoginEncryptResponse.VerifyToken = make([]byte, verifyTokenLength)
	_, err = reader.Read(packetServerLoginEncryptResponse.VerifyToken)
	if err != nil {
		return
	}
	decode = packetServerLoginEncryptResponse
	return
}

func (this *packetServerLoginEncryptResponseCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerLoginEncryptResponse := encode.(*PacketServerLoginEncryptResponse)
	err = packet.WriteVarInt(writer, util, len(packetServerLoginEncryptResponse.SharedSecret))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, util, len(packetServerLoginEncryptResponse.VerifyToken))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.VerifyToken)
	return
}
