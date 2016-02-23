package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientLoginSetCompression struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientLoginSetCompression) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientLoginSetCompression := new(minecraft.PacketClientLoginSetCompression)
	packetClientLoginSetCompression.IdFrom(this.IdMap)
	packetClientLoginSetCompression.Threshold, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	decode = packetClientLoginSetCompression
	return
}

func (this *CodecClientLoginSetCompression) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientLoginSetCompression := encode.(*minecraft.PacketClientLoginSetCompression)
	err = packet.WriteVarInt(writer, packetClientLoginSetCompression.Threshold)
	return
}
