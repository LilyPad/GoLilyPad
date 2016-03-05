package packet

import (
	"bytes"
	"github.com/klauspost/compress/zlib"
	"io"
)

type ZlibToggleReader struct {
	bytes         []byte
	rawReader     io.Reader
	zlibReader    io.ReadCloser
	compression   bool
	currentReader io.Reader
}

func NewZlibToggleReaderBuffer(rawBytes []byte, zlibBytes []byte) (this *ZlibToggleReader, err error) {
	this = new(ZlibToggleReader)
	this.rawReader = bytes.NewBuffer(rawBytes)
	this.zlibReader, err = zlib.NewReader(bytes.NewBuffer(zlibBytes))
	if err != nil {
		return
	}
	this.SetCompression(true)
	return
}

func (this *ZlibToggleReader) Read(p []byte) (n int, err error) {
	n, err = io.ReadFull(this.currentReader, p)
	if err == io.ErrUnexpectedEOF {
		err = nil
	}
	return
}

func (this *ZlibToggleReader) SetRaw(raw bool) {
	this.SetCompression(!raw)
}

func (this *ZlibToggleReader) SetCompression(compression bool) {
	this.compression = compression
	if this.compression {
		this.currentReader = this.zlibReader
	} else {
		this.currentReader = this.rawReader
	}
}

func (this *ZlibToggleReader) Close() (err error) {
	err = this.zlibReader.Close()
	return
}
