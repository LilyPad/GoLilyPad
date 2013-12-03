package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestRedirect struct {
	Server string
	Player string
}

func (this *RequestRedirect) Id() int {
	return REQUEST_REDIRECT
}

type RequestRedirectCodec struct {
	
}

func (this *RequestRedirectCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestRedirect := &RequestRedirect{}
	requestRedirect.Server, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestRedirect.Player, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return requestRedirect, nil
}

func (this *RequestRedirectCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
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

func (this *ResultRedirect) Id() int {
	return REQUEST_REDIRECT
}

type ResultRedirectCodec struct {

}

func (this *ResultRedirectCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	return &ResultRedirect{}, nil
}

func (this *ResultRedirectCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}