package packet

import "io"

type FullReader struct {
	reader io.Reader
}

func (fullReader *FullReader) Read(p []byte) (n int, err error) {
	return io.ReadFull(fullReader.reader, p)
}
