package packet

import "io"
import "errors"

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
		err = errors.New("Packet Id is below zero")
		return
	}
	if id >= len(this.packetCodecs) {
		err = errors.New("Packet Id is above maximum")
		return
	} 
	codec := this.packetCodecs[id]
	if codec == nil {
		err = errors.New("Packet Id does not have a codec")
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
		err = errors.New("Packet Id is below zero")
		return
	}
	if id >= len(this.packetCodecs) {
		err = errors.New("Packet Id is above maximum")
		return
	}
	codec := this.packetCodecs[id]
	if codec == nil {
		err = errors.New("Packet Id does not have a codec")
		return
	}
	return codec.Encode(writer, util, packet)
}
