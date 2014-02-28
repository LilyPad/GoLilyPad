package packet

import "io"

type fullReader struct {
	reader io.Reader
}

func NewFullReader(reader io.Reader) io.Reader {
	return &fullReader{reader}
}

func (this *fullReader) Read(p []byte) (n int, err error) {
	return io.ReadFull(this.reader, p)
}
