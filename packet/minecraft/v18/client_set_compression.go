package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientSetCompression struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientSetCompression) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientSetCompression := new(minecraft.PacketClientSetCompression)
	packetClientSetCompression.IdFrom(this.IdMap)
	packetClientSetCompression.Threshold, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	decode = packetClientSetCompression
	return
}

func (this *CodecClientSetCompression) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientSetCompression := encode.(*minecraft.PacketClientSetCompression)
	err = packet.WriteVarInt(writer, packetClientSetCompression.Threshold)
	return
}
