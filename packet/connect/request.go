package connect

import "io"

type Request interface {
	Id() int
}

type RequestCodec interface {
	Decode(reader io.Reader, util []byte) (request Request, err error)
	Encode(writer io.Writer, util []byte, request Request) (err error)
}

type Result interface {
	Id() int
}

type ResultCodec interface {
	Decode(reader io.Reader, util []byte) (result Result, err error)
	Encode(writer io.Writer, util []byte, result Result) (err error)
}