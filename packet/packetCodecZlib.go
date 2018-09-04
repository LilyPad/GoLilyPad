package packet

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/klauspost/compress/zlib"
	"io"
)

type PacketCodecZlib struct {
	codec      PacketCodec
	threshold  int
	level      int
	zlibWriter *zlib.Writer
}

func NewPacketCodecZlib(threshold int) (this *PacketCodecZlib) {
	this = NewPacketCodecZlibLevel(threshold, zlib.DefaultCompression)
	return
}

func NewPacketCodecZlibLevel(threshold int, level int) (this *PacketCodecZlib) {
	this = new(PacketCodecZlib)
	this.threshold = threshold
	this.level = level
	this.zlibWriter, _ = zlib.NewWriterLevel(nil, level)
	return
}

func (this *PacketCodecZlib) Decode(reader io.Reader) (packet Packet, err error) {
	rawBytes := reader.(Byteser).Bytes() // FIXME assuming the caller is a Byteser is a bad idea
	length, err := ReadVarInt(reader)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Decode, Compressed length is below zero: %d", length))
		return
	}
	if length == 0 {
		packet, err = this.codec.Decode(reader)
	} else {
		zlibBytes := reader.(Byteser).Bytes() // FIXME assuming the caller is a Byteser is a bad idea
		var zlibReader io.ReadCloser
		zlibReader, err = NewZlibToggleReaderBuffer(rawBytes, zlibBytes)
		if err != nil {
			return
		}
		packet, err = this.codec.Decode(zlibReader)
		if err != nil {
			return
		}
		zlibReader.Close()
	}
	return
}

func (this *PacketCodecZlib) Encode(writer io.Writer, packet Packet) (err error) {
	buffer := new(bytes.Buffer)
	err = this.codec.Encode(buffer, packet)
	if err != nil {
		return
	}
	if raw, ok := packet.(PacketRaw); ok && raw.Raw() {
		_, err = buffer.WriteTo(writer)
	} else if buffer.Len() < this.threshold {
		err = WriteVarInt(writer, 0)
		if err != nil {
			return
		}
		_, err = buffer.WriteTo(writer)
	} else {
		err = WriteVarInt(writer, buffer.Len())
		if err != nil {
			return
		}
		this.zlibWriter.Reset(writer)
		_, err = buffer.WriteTo(this.zlibWriter)
		if err != nil {
			return
		}
		err = this.zlibWriter.Flush()
	}
	return
}

func (this *PacketCodecZlib) SetCodec(codec PacketCodec) {
	this.codec = codec
}
