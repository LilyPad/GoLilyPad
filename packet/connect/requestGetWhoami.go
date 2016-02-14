package connect

import (
	"io"
	"github.com/suedadam/GoLilyPad/packet"
)

type RequestGetWhoami struct {

}

func NewRequestGetWhoami() (this *RequestGetWhoami) {
	this = new(RequestGetWhoami)
	return
}

func (this *RequestGetWhoami) Id() int {
	return REQUEST_GET_WHOAMI
}

type requestGetWhoamiCodec struct {

}

func (this *requestGetWhoamiCodec) Decode(reader io.Reader) (request Request, err error) {
	request = new(RequestGetWhoami)
	return
}

func (this *requestGetWhoamiCodec) Encode(writer io.Writer, request Request) (err error) {
	return
}

type ResultGetWhoami struct {
	Whoiam string
}

func NewResultGetWhoami(whoiam string) (this *ResultGetWhoami) {
	this = new(ResultGetWhoami)
	this.Whoiam = whoiam
	return
}

func (this *ResultGetWhoami) Id() int {
	return REQUEST_GET_WHOAMI
}

type resultGetWhoamiCodec struct {

}

func (this *resultGetWhoamiCodec) Decode(reader io.Reader) (result Result, err error) {
	resultGetWhoami := new(ResultGetWhoami)
	resultGetWhoami.Whoiam, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	result = resultGetWhoami
	return
}

func (this *resultGetWhoamiCodec) Encode(writer io.Writer, result Result) (err error) {
	err = packet.WriteString(writer, result.(*ResultGetWhoami).Whoiam)
	return
}
