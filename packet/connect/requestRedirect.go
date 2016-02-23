package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type RequestRedirect struct {
	Server string
	Player string
}

func NewRequestRedirect(server string, player string) (this *RequestRedirect) {
	this = new(RequestRedirect)
	this.Server = server
	this.Player = player
	return
}

func (this *RequestRedirect) Id() int {
	return REQUEST_REDIRECT
}

type requestRedirectCodec struct {
}

func (this *requestRedirectCodec) Decode(reader io.Reader) (request Request, err error) {
	requestRedirect := new(RequestRedirect)
	requestRedirect.Server, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	requestRedirect.Player, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	request = requestRedirect
	return
}

func (this *requestRedirectCodec) Encode(writer io.Writer, request Request) (err error) {
	requestRedirect := request.(*RequestRedirect)
	err = packet.WriteString(writer, requestRedirect.Server)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, requestRedirect.Player)
	return
}

type ResultRedirect struct {
}

func NewResultRedirect() (this *ResultRedirect) {
	this = new(ResultRedirect)
	return
}

func (this *ResultRedirect) Id() int {
	return REQUEST_REDIRECT
}

type resultRedirectCodec struct {
}

func (this *resultRedirectCodec) Decode(reader io.Reader) (result Result, err error) {
	result = new(ResultRedirect)
	return
}

func (this *resultRedirectCodec) Encode(writer io.Writer, result Result) (err error) {
	return
}
