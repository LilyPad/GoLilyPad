package packet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type PacketCodecRegistry struct {
	EncodeCodecs    []PacketCodec
	DecodeCodecs    []PacketCodec
	interceptDecode PacketIntercept
	interceptEncode PacketIntercept
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
	var buffer io.Reader
	var bufferPayload []byte
	if this.interceptDecode != nil {
		bufferPayload, err = ioutil.ReadAll(reader)
		if err != nil {
			return
		}
		buffer = bytes.NewBuffer(bufferPayload)
	} else {
		buffer = reader
	}
	id, err := ReadVarInt(buffer)
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
	packet, err = codec.Decode(buffer)
	if err != nil {
		err = PacketDecodeError{id, codec, err}
		return
	}
	if this.interceptDecode != nil {
		err = this.interceptDecode(packet, bytes.NewBuffer(bufferPayload))
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
	buffer := new(bytes.Buffer)
	if raw, ok := packet.(PacketRaw); ok && raw.Raw() {
		err = codec.Encode(buffer, packet)
	} else {
		err = WriteVarInt(buffer, id)
		if err != nil {
			return
		}
		err = codec.Encode(buffer, packet)
	}
	if err != nil {
		return
	}
	if this.interceptEncode != nil {
		err = this.interceptEncode(packet, buffer)
		if err != nil {
			return
		}
	}
	_, err = buffer.WriteTo(writer)
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

func (this *PacketCodecRegistry) SetInterceptDecode(intercept PacketIntercept) {
	this.interceptDecode = intercept
}

func (this *PacketCodecRegistry) SetInterceptEncode(intercept PacketIntercept) {
	this.interceptEncode = intercept
}
