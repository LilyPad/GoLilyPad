package packet

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
)

type PacketCodecZlib struct {
	codec PacketCodec
	threshold int
	level int
}

func NewPacketCodecZlib(threshold int) (this *PacketCodecZlib) {
	this = NewPacketCodecZlibLevel(threshold, zlib.DefaultCompression)
	return
}

func NewPacketCodecZlibLevel(threshold int, level int) (this *PacketCodecZlib) {
	this = new(PacketCodecZlib)
	this.threshold = threshold
	this.level = level
	return
}

func (this *PacketCodecZlib) Decode(reader io.Reader, util []byte) (packet Packet, err error) {
	rawBytes := reader.(Byteser).Bytes() // FIXME assuming the caller is a Byteser is a bad idea
	length, err := ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Decode, Compressed length is below zero: %d", length))
		return
	}
	if length == 0 {
		packet, err = this.codec.Decode(reader, util)
	} else {
		zlibBytes := reader.(Byteser).Bytes() // FIXME assuming the caller is a Byteser is a bad idea
		var zlibReader io.ReadCloser
		zlibReader, err = NewZlibToggleReaderBuffer(rawBytes, zlibBytes)
		if err != nil {
			return
		}
		packet, err = this.codec.Decode(zlibReader, util)
		if err != nil {
			return
		}
		err = zlibReader.Close()
	}
	return
}

func (this *PacketCodecZlib) Encode(writer io.Writer, util []byte, packet Packet) (err error) {
	buffer := new(bytes.Buffer)
	err = this.codec.Encode(buffer, util, packet)
	if err != nil {
		return
	}
	if raw, ok := packet.(PacketRaw); ok && raw.Raw() {
		_, err = buffer.WriteTo(writer)
	} else if buffer.Len() < this.threshold {
		err = WriteVarInt(writer, util, 0)
		if err != nil {
			return
		}
		_, err = buffer.WriteTo(writer)
	} else {
		err = WriteVarInt(writer, util, buffer.Len())
		if err != nil {
			return
		}
		var zlibWriter io.WriteCloser
		zlibWriter, err = zlib.NewWriterLevel(writer, this.level)
		if err != nil {
			return
		}
		_, err = buffer.WriteTo(zlibWriter)
		if err != nil {
			return
		}
		err = zlibWriter.Close()
	}
	return
}

func (this *PacketCodecZlib) SetCodec(codec PacketCodec) {
	this.codec = codec
}
