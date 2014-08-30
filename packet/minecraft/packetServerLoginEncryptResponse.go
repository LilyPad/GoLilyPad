package minecraft

import (
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
	sharedSecretSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.SharedSecret = make([]byte, sharedSecretSize)
	_, err = reader.Read(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	verifyTokenSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.VerifyToken = make([]byte, verifyTokenSize)
	_, err = reader.Read(packetServerLoginEncryptResponse.VerifyToken)
	if err != nil {
		return
	}
	decode = packetServerLoginEncryptResponse
	return
}

func (this *packetServerLoginEncryptResponseCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerLoginEncryptResponse := encode.(*PacketServerLoginEncryptResponse)
	err = packet.WriteUint16(writer, util, uint16(len(packetServerLoginEncryptResponse.SharedSecret)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetServerLoginEncryptResponse.VerifyToken)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.VerifyToken)
	return
}
