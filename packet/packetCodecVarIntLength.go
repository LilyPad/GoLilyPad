package packet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type PacketCodecVarIntLength struct {
	packetCodec PacketCodec
}

func NewPacketCodecVarIntLength(packetCodec PacketCodec) (this *PacketCodecVarIntLength) {
	this = new(PacketCodecVarIntLength)
	this.packetCodec = packetCodec
	return
}

func (this *PacketCodecVarIntLength) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	length, err := ReadVarInt(reader, util)
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
	packet, err = this.packetCodec.Decode(bytes.NewReader(payload), util)
	return
}

func (this *PacketCodecVarIntLength) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	buffer := new(bytes.Buffer)
	err = this.packetCodec.Encode(buffer, util, packet)
	if err != nil {
		return
	}
	err = WriteVarInt(writer, util, buffer.Len())
	if err != nil {
		return
	}
	_, err = buffer.WriteTo(writer)
	return
}
