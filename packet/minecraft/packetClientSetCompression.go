package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type PacketClientSetCompression struct {
	Threshold int
}

func NewPacketClientSetCompression(threshold int) (this *PacketClientSetCompression) {
	this = new(PacketClientSetCompression)
	this.Threshold = threshold
	return
}

func (this *PacketClientSetCompression) Id() int {
	return PACKET_CLIENT_SET_COMPRESSION
}

type packetClientSetCompressionCodec struct {
}

func (this *packetClientSetCompressionCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientSetCompression := new(PacketClientSetCompression)
	packetClientSetCompression.Threshold, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	decode = packetClientSetCompression
	return
}

func (this *packetClientSetCompressionCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientSetCompression := encode.(*PacketClientSetCompression)
	err = packet.WriteVarInt(writer, packetClientSetCompression.Threshold)
	return
}
