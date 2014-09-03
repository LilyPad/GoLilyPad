package packet

import (
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"
)

type ZlibToggleReader struct {
	bytes []byte
	rawReader io.Reader
	zlibReader io.Reader
	compression bool
	buffer bool
	buffered bool
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
	if this.buffer && this.compression {
		var uncompressed []byte
		uncompressed, err = ioutil.ReadAll(this.zlibReader)
		if err != nil {
			return
		}
		err = this.zlibReader.(io.ReadCloser).Close()
		if err != nil {
			return
		}
		this.zlibReader = bytes.NewBuffer(uncompressed)
		if this.compression {
			this.currentReader = this.zlibReader
		}
		this.buffer = false
		this.buffered = true
	}
	n, err = this.currentReader.Read(p)
	return
}

func (this *ZlibToggleReader) Buffer() {
	if this.buffered {
		return
	}
	this.buffer = true
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
	if zlibReader, ok := this.zlibReader.(io.ReadCloser); ok {
		zlibReader.Close()
	}
	return
}
