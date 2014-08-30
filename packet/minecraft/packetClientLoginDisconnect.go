package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientLoginDisconnect struct {
	Json string
}

func NewPacketClientLoginDisconnect(json string) (this *PacketClientLoginDisconnect) {
	this = new(PacketClientLoginDisconnect)
	this.Json = json
	return
}

func (this *PacketClientLoginDisconnect) Id() int {
	return PACKET_CLIENT_LOGIN_DISCONNECT
}

type packetClientLoginDisconnectCodec struct {

}

func (this *packetClientLoginDisconnectCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginDisconnect := new(PacketClientLoginDisconnect)
	packetClientLoginDisconnect.Json, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	decode = packetClientLoginDisconnect
	return
}

func (this *packetClientLoginDisconnectCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginDisconnect := encode.(*PacketClientLoginDisconnect)
	err = packet.WriteString(writer, util, packetClientLoginDisconnect.Json)
	return
}
