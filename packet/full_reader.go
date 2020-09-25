package packet

import (
	"io"
)

type fullReader struct {
	Reader io.Reader
}

func NewFullReader(reader io.Reader) (this io.Reader) {
	fullReader := new(fullReader)
	fullReader.Reader = reader
	this = fullReader
	return
}

func (this *fullReader) Read(p []byte) (n int, err error) {
	n, err = io.ReadFull(this.Reader, p)
	return
}
