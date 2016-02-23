package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientDisconnect struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientDisconnect) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientDisconnect := new(minecraft.PacketClientDisconnect)
	packetClientDisconnect.IdFrom(this.IdMap)
	packetClientDisconnect.Json, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetClientDisconnect
	return
}

func (this *CodecClientDisconnect) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientDisconnect := encode.(*minecraft.PacketClientDisconnect)
	err = packet.WriteString(writer, packetClientDisconnect.Json)
	return
}
