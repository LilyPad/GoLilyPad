package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
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

func (this *requestRedirectCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestRedirect := new(RequestRedirect)
	requestRedirect.Server, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestRedirect.Player, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	request = requestRedirect
	return
}

func (this *requestRedirectCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	requestRedirect := request.(*RequestRedirect)
	err = packet.WriteString(writer, util, requestRedirect.Server)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, requestRedirect.Player)
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

func (this *resultRedirectCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	result = new(ResultRedirect)
	return
}

func (this *resultRedirectCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}
