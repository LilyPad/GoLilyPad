package packet

import (
	"errors"
	"fmt"
	"io"
)

type PacketCodecRegistry struct {
	EncodeCodecs []PacketCodec
	DecodeCodecs []PacketCodec
}

type PacketDecodeError struct {
	Id    int
	Codec PacketCodec
	Err   error
}

func (this PacketDecodeError) Error() string {
	return fmt.Sprintf("Error decoding id: %d codec: \"%s\" err: \"%s\"", this.Id, this.Codec, this.Err)
}

func NewPacketCodecRegistry(codecs []PacketCodec) (this *PacketCodecRegistry) {
	this = new(PacketCodecRegistry)
	this.EncodeCodecs = codecs
	this.DecodeCodecs = codecs
	return
}

func NewPacketCodecRegistryDual(encodeCodecs []PacketCodec, decodeCodecs []PacketCodec) (this *PacketCodecRegistry) {
	this = new(PacketCodecRegistry)
	this.EncodeCodecs = encodeCodecs
	this.DecodeCodecs = decodeCodecs
	return
}

func (this *PacketCodecRegistry) Decode(reader io.Reader) (packet Packet, err error) {
	id, err := ReadVarInt(reader)
	if err != nil {
		return
	}
	if id < 0 {
		err = errors.New(fmt.Sprintf("Decode, Packet Id is below zero: %d", id))
		return
	}
	if id >= len(this.DecodeCodecs) {
		err = errors.New(fmt.Sprintf("Decode, Packet Id is above maximum: %d", id))
		return
	}
	codec := this.DecodeCodecs[id]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Decode, Packet Id does not have a codec: %d", id))
		return
	}
	packet, err = codec.Decode(reader)
	if err != nil {
		err = PacketDecodeError{id, codec, err}
	}
	return
}

func (this *PacketCodecRegistry) Encode(writer io.Writer, packet Packet) (err error) {
	id := packet.Id()
	if id < 0 {
		err = errors.New(fmt.Sprintf("Encode, Packet Id is below zero: %d", id))
		return
	}
	if id >= len(this.EncodeCodecs) {
		err = errors.New(fmt.Sprintf("Encode, Packet Id is above maximum: %d", id))
		return
	}
	codec := this.EncodeCodecs[id]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Encode, Packet Id does not have a codec: %d", id))
		return
	}
	if raw, ok := packet.(PacketRaw); ok && raw.Raw() {
		err = codec.Encode(writer, packet)
	} else {
		err = WriteVarInt(writer, id)
		if err != nil {
			return
		}
		err = codec.Encode(writer, packet)
	}
	return
}

func (this *PacketCodecRegistry) Flip() (thisCopy *PacketCodecRegistry) {
	thisCopy = this.Copy()
	encodeCodecs := thisCopy.EncodeCodecs
	thisCopy.EncodeCodecs = thisCopy.DecodeCodecs
	thisCopy.DecodeCodecs = encodeCodecs
	return
}

func (this *PacketCodecRegistry) Copy() (thisCopy *PacketCodecRegistry) {
	thisCopy = new(PacketCodecRegistry)
	thisCopy.EncodeCodecs = make([]PacketCodec, len(this.EncodeCodecs))
	copy(thisCopy.EncodeCodecs, this.EncodeCodecs)
	thisCopy.DecodeCodecs = make([]PacketCodec, len(this.DecodeCodecs))
	copy(thisCopy.DecodeCodecs, this.DecodeCodecs)
	return
}

func (this *PacketCodecRegistry) SetCodec(codec PacketCodec) {
	panic("PacketCodecRegistry must be last in the pipeline")
}
