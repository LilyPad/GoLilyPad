package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type packetServerLoginEncryptResponseCodec17 struct {

}

func (this *packetServerLoginEncryptResponseCodec17) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerLoginEncryptResponse := new(PacketServerLoginEncryptResponse)
	sharedSecretLength, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.SharedSecret = make([]byte, sharedSecretLength)
	_, err = reader.Read(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	verifyTokenLength, err := packet.ReadUint16(reader, util)
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

func (this *packetServerLoginEncryptResponseCodec17) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
