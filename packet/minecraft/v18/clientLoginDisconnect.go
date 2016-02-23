package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientLoginDisconnect struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientLoginDisconnect) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientLoginDisconnect := new(minecraft.PacketClientLoginDisconnect)
	packetClientLoginDisconnect.IdFrom(this.IdMap)
	packetClientLoginDisconnect.Json, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetClientLoginDisconnect
	return
}

func (this *CodecClientLoginDisconnect) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientLoginDisconnect := encode.(*minecraft.PacketClientLoginDisconnect)
	err = packet.WriteString(writer, packetClientLoginDisconnect.Json)
	return
}
