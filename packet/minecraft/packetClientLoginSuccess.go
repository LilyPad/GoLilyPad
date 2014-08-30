package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientLoginSuccess struct {
	UUID string
	Name string
}

func NewPacketClientLoginSuccess(uuid string, name string) (this *PacketClientLoginSuccess) {
	this = new(PacketClientLoginSuccess)
	this.UUID = uuid
	this.Name = name
	return
}

func (this *PacketClientLoginSuccess) Id() int {
	return PACKET_CLIENT_LOGIN_SUCCESS
}

type packetClientLoginSuccessCodec struct {

}

func (this *packetClientLoginSuccessCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginSuccess := new(PacketClientLoginSuccess)
	packetClientLoginSuccess.UUID, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetClientLoginSuccess.Name, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	decode = packetClientLoginSuccess
	return
}

func (this *packetClientLoginSuccessCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginSuccess := encode.(*PacketClientLoginSuccess)
	err = packet.WriteString(writer, util, packetClientLoginSuccess.UUID)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetClientLoginSuccess.Name)
	return
}
