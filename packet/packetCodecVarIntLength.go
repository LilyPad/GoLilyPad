package packet

import "bytes"
import "errors"
import "fmt"
import "io"

type PacketCodecVarIntLength struct {
	packetCodec PacketCodec
}

func NewPacketCodecVarIntLength(packetCodec PacketCodec) *PacketCodecVarIntLength {
	return &PacketCodecVarIntLength{packetCodec}
}

func (this *PacketCodecVarIntLength) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	length, err := ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Packet length is negative: %d", length))
		return
	}
	if length > 2097151 { // 2^21
		err = errors.New(fmt.Sprintf("Packet length is above maximum: %d", length))
		return
	}
	payload := make([]byte, length)
	_, err = reader.Read(payload)
	if err != nil {
		return
	}
	return this.packetCodec.Decode(bytes.NewReader(payload), util)
}

func (this *PacketCodecVarIntLength) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	buffer := &bytes.Buffer{}
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
