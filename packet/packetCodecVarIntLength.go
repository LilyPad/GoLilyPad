package packet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type PacketCodecVarIntLength struct {
	codec           PacketCodec
	interceptDecode PacketIntercept
	interceptEncode PacketIntercept
}

func NewPacketCodecVarIntLength() (this *PacketCodecVarIntLength) {
	this = new(PacketCodecVarIntLength)
	return
}

func (this *PacketCodecVarIntLength) Decode(reader io.Reader) (packet Packet, err error) {
	length, err := ReadVarInt(reader)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Decode, Packet length is below zero: %d", length))
		return
	}
	if length > 1048576 { // 2^(21-1)
		err = errors.New(fmt.Sprintf("Decode, Packet length is above maximum: %d", length))
		return
	}
	payload := make([]byte, length)
	_, err = reader.Read(payload)
	if err != nil {
		return
	}
	packet, err = this.codec.Decode(bytes.NewBuffer(payload))
	if this.interceptDecode != nil {
		if err != nil {
			return
		}
		err = this.interceptDecode(packet, bytes.NewBuffer(payload))
	}
	return
}

func (this *PacketCodecVarIntLength) Encode(writer io.Writer, packet Packet) (err error) {
	buffer := new(bytes.Buffer)
	err = this.codec.Encode(buffer, packet)
	if this.interceptEncode != nil {
		if err != nil {
			return
		}
		err = this.interceptEncode(packet, buffer)
	}
	if err != nil {
		return
	}
	err = WriteVarInt(writer, buffer.Len())
	if err != nil {
		return
	}
	_, err = buffer.WriteTo(writer)
	return
}

func (this *PacketCodecVarIntLength) SetCodec(codec PacketCodec) {
	this.codec = codec
}

func (this *PacketCodecVarIntLength) SetInterceptDecode(intercept PacketIntercept) {
	this.interceptDecode = intercept
}

func (this *PacketCodecVarIntLength) SetInterceptEncode(intercept PacketIntercept) {
	this.interceptEncode = intercept
}
