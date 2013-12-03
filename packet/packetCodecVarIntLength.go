package packet

import "bytes"
import "errors"
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
		err = errors.New("Packet length is negative")
		return
	}
	payload := make([]byte, length)
	_, err = reader.Read(payload)
	if err != nil {
		return
	}
	buffer := bytes.NewBuffer(payload)
	return this.packetCodec.Decode(buffer, util)
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
