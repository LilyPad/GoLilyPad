package connect

import (
	"io"
)

type Request interface {
	Id() int
}

type RequestCodec interface {
	Decode(reader io.Reader) (request Request, err error)
	Encode(writer io.Writer, request Request) (err error)
}

type Result interface {
	Id() int
}

type ResultCodec interface {
	Decode(reader io.Reader) (result Result, err error)
	Encode(writer io.Writer, result Result) (err error)
}
