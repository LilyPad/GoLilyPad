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
		var zlibReader io.ReadCloser
		zlibReader, err = zlib.NewReader(reader)
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
	if buffer.Len() >= this.threshold {
		err = WriteVarInt(writer, util, buffer.Len())
		if err != nil {
			return
		}
		var zlibWriter *zlib.Writer
		zlibWriter, err = zlib.NewWriterLevel(writer, this.level)
		if err != nil {
			return
		}
		_, err = buffer.WriteTo(zlibWriter)
		if err != nil {
			return
		}
		err = zlibWriter.Close()
	} else {
		err = WriteVarInt(writer, util, 0)
		if err != nil {
			return
		}
		_, err = buffer.WriteTo(writer)
	}
	return
}

func (this *PacketCodecZlib) SetCodec(codec PacketCodec) {
	this.codec = codec
}
