package packet

import "io"

type PacketCodecVariable struct {
	encodePacketCodec PacketCodec
	decodePacketCodec PacketCodec
}

func NewPacketCodecVariable(encodePacketCodec PacketCodec, decodePacketCodec PacketCodec) *PacketCodecVariable {
	return &PacketCodecVariable{encodePacketCodec, decodePacketCodec}
}

func (this *PacketCodecVariable) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	return this.decodePacketCodec.Decode(reader, util)
}

func (this *PacketCodecVariable) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	return this.encodePacketCodec.Encode(writer, util, packet)
}

func (this *PacketCodecVariable) SetDecodeCodec(packetCodec PacketCodec) {
	this.decodePacketCodec = packetCodec
}

func (this *PacketCodecVariable) DecodeCodec() PacketCodec {
	return this.decodePacketCodec
}

func (this *PacketCodecVariable) SetEncodeCodec(packetCodec PacketCodec) {
	this.encodePacketCodec = packetCodec
}

func (this *PacketCodecVariable) EncodeCodec() PacketCodec {
	return this.encodePacketCodec
}
