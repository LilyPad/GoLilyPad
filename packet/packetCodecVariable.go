package packet

import (
	"io"
)

type PacketCodecVariable struct {
	encodePacketCodec PacketCodec
	decodePacketCodec PacketCodec
}

func NewPacketCodecVariable(encodePacketCodec PacketCodec, decodePacketCodec PacketCodec) (this *PacketCodecVariable) {
	this = new(PacketCodecVariable)
	this.encodePacketCodec = encodePacketCodec
	this.decodePacketCodec = decodePacketCodec
	return
}

func (this *PacketCodecVariable) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	packet, err = this.decodePacketCodec.Decode(reader, util)
	return
}

func (this *PacketCodecVariable) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	err = this.encodePacketCodec.Encode(writer, util, packet)
	return
}

func (this *PacketCodecVariable) SetDecodeCodec(packetCodec PacketCodec) {
	this.decodePacketCodec = packetCodec
}

func (this *PacketCodecVariable) DecodeCodec() (packetCodec PacketCodec) {
	packetCodec = this.decodePacketCodec
	return
}

func (this *PacketCodecVariable) SetEncodeCodec(packetCodec PacketCodec) {
	this.encodePacketCodec = packetCodec
}

func (this *PacketCodecVariable) EncodeCodec() (packetCodec PacketCodec) {
	packetCodec = this.encodePacketCodec
	return
}
