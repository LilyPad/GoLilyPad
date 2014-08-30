package packet

import (
	"errors"
	"fmt"
	"io"
)

type PacketCodecRegistry struct {
	packetCodecs []PacketCodec
}

func NewPacketCodecRegistry(packetCodecs []PacketCodec) (this *PacketCodecRegistry) {
	this = new(PacketCodecRegistry)
	this.packetCodecs = packetCodecs
	return
}

func (this *PacketCodecRegistry) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	id, err := ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if id < 0 {
		err = errors.New(fmt.Sprintf("Decode, Packet Id is below zero: %d", id))
		return
	}
	if id >= len(this.packetCodecs) {
		err = errors.New(fmt.Sprintf("Decode, Packet Id is above maximum: %d", id))
		return
	}
	codec := this.packetCodecs[id]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Decode, Packet Id does not have a codec: %d", id))
		return
	}
	packet, err = codec.Decode(reader, util)
	return
}

func (this *PacketCodecRegistry) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	id := packet.Id()
	err = WriteVarInt(writer, util, id)
	if err != nil {
		return
	}
	if id < 0 {
		err = errors.New(fmt.Sprintf("Encode, Packet Id is below zero: %d", id))
		return
	}
	if id >= len(this.packetCodecs) {
		err = errors.New(fmt.Sprintf("Encode, Packet Id is above maximum: %d", id))
		return
	}
	codec := this.packetCodecs[id]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Encode, Packet Id does not have a codec: %d", id))
		return
	}
	err = codec.Encode(writer, util, packet)
	return
}
