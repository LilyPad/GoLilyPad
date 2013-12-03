package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestGetWhoami struct {

}

func (this *RequestGetWhoami) Id() int {
	return REQUEST_GET_WHOAMI
}

type RequestGetWhoamiCodec struct {

}

func (this *RequestGetWhoamiCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	return &RequestGetWhoami{}, nil
}

func (this *RequestGetWhoamiCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	return
}

type ResultGetWhoami struct {
	Whoiam string
}

func (this *ResultGetWhoami) Id() int {
	return REQUEST_GET_WHOAMI
}

type ResultGetWhoamiCodec struct {
	
}

func (this *ResultGetWhoamiCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultGetWhoami := &ResultGetWhoami{}
	resultGetWhoami.Whoiam, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return resultGetWhoami, nil
}

func (this *ResultGetWhoamiCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	err = packet.WriteString(writer, util, result.(*ResultGetWhoami).Whoiam)
	return
}
