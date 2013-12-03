package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestAuthenticate struct {
	Username string
	Password string
}

func (this *RequestAuthenticate) Id() int {
	return REQUEST_AUTHENTICATE
}

type RequestAuthenticateCodec struct {
	
}

func (this *RequestAuthenticateCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestAuthenticate := &RequestAuthenticate{}
	requestAuthenticate.Username, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestAuthenticate.Password, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return requestAuthenticate, nil
}

func (this *RequestAuthenticateCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	requestAuthenticate := request.(*RequestAuthenticate)
	err = packet.WriteString(writer, util, requestAuthenticate.Username)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, requestAuthenticate.Password)
	return
}

type ResultAuthenticate struct {

}

func (this *ResultAuthenticate) Id() int {
	return REQUEST_AUTHENTICATE
}

type ResultAuthenticateCodec struct {

}

func (this *ResultAuthenticateCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	return &ResultAuthenticate{}, nil
}

func (this *ResultAuthenticateCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}