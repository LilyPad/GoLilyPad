package packet

import "errors"
import "fmt"
import "io"

type PacketCodecRegistry struct {
	packetCodecs []PacketCodec	
}

func NewPacketCodecRegistry(packetCodecs []PacketCodec) *PacketCodecRegistry {
	return &PacketCodecRegistry{packetCodecs}
}

func (this *PacketCodecRegistry) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	id, err := ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if id < 0 {
		err = errors.New(fmt.Sprintf("Packet Id is below zero: %d", id))
		return
	}
	if id >= len(this.packetCodecs) {
		err = errors.New(fmt.Sprintf("Packet Id is above maximum: %d", id))
		return
	} 
	codec := this.packetCodecs[id]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Packet Id does not have a codec: %d", id))
		return
	}
	return codec.Decode(reader, util)
}

func (this *PacketCodecRegistry) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	id := packet.Id()
	err = WriteVarInt(writer, util, id)
	if err != nil {
		return
	}
	if id < 0 {
		err = errors.New(fmt.Sprintf("Packet Id is below zero: %d", id))
		return
	}
	if id >= len(this.packetCodecs) {
		err = errors.New(fmt.Sprintf("Packet Id is above maximum: %d", id))
		return
	}
	codec := this.packetCodecs[id]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Packet Id does not have a codec: %d", id))
		return
	}
	return codec.Encode(writer, util, packet)
}
