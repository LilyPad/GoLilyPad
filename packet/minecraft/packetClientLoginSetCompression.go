package minecraft

import (
	"io"
	"github.com/suedadam/GoLilyPad/packet"
)

type PacketClientLoginSetCompression struct {
	Threshold int
}

func NewPacketClientLoginSetCompression(threshold int) (this *PacketClientLoginSetCompression) {
	this = new(PacketClientLoginSetCompression)
	this.Threshold = threshold
	return
}

func (this *PacketClientLoginSetCompression) Id() int {
	return PACKET_CLIENT_LOGIN_SET_COMPRESSION
}

type packetClientLoginSetCompressionCodec struct {

}

func (this *packetClientLoginSetCompressionCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientLoginSetCompression := new(PacketClientLoginSetCompression)
	packetClientLoginSetCompression.Threshold, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	decode = packetClientLoginSetCompression
	return
}

func (this *packetClientLoginSetCompressionCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientLoginSetCompression := encode.(*PacketClientLoginSetCompression)
	err = packet.WriteVarInt(writer, packetClientLoginSetCompression.Threshold)
	return
}
