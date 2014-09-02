package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type packetClientLoginEncryptRequestCodec17 struct {

}

func (this *packetClientLoginEncryptRequestCodec17) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginEncryptRequest := new(PacketClientLoginEncryptRequest)
	packetClientLoginEncryptRequest.ServerId, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	publicKeyLength, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetClientLoginEncryptRequest.PublicKey = make([]byte, publicKeyLength)
	_, err = reader.Read(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	verifyTokenLength, err := packet.ReadUint16(reader, util)
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

func (this *packetClientLoginEncryptRequestCodec17) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginEncryptRequest := encode.(*PacketClientLoginEncryptRequest)
	err = packet.WriteString(writer, util, packetClientLoginEncryptRequest.ServerId)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetClientLoginEncryptRequest.PublicKey)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetClientLoginEncryptRequest.VerifyToken)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.VerifyToken)
	return
}
