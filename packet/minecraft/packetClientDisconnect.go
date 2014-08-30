package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientDisconnect struct {
	Json string
}

func NewPacketClientDisconnect(json string) (this *PacketClientDisconnect) {
	this = new(PacketClientDisconnect)
	this.Json = json
	return
}

func (this *PacketClientDisconnect) Id() int {
	return PACKET_CLIENT_DISCONNECT
}

type packetClientDisconnectCodec struct {

}

func (this *packetClientDisconnectCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientDisconnect := new(PacketClientDisconnect)
	packetClientDisconnect.Json, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	decode = packetClientDisconnect
	return
}

func (this *packetClientDisconnectCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientDisconnect := encode.(*PacketClientDisconnect)
	err = packet.WriteString(writer, util, packetClientDisconnect.Json)
	return
}
