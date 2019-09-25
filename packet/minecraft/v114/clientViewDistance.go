package v114

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientViewDistance struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientViewDistance) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientViewDistance := new(minecraft.PacketClientViewDistance)
	packetClientViewDistance.IdFrom(this.IdMap)
	packetClientViewDistance.ViewDistance, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	decode = packetClientViewDistance
	return
}

func (this *CodecClientViewDistance) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientViewDistance := encode.(*minecraft.PacketClientViewDistance)
	err = packet.WriteVarInt(writer, packetClientViewDistance.ViewDistance)
	return
}
