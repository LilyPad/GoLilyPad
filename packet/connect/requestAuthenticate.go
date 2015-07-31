package connect

import "io"
import "github.com/suedadam/GoLilyPad/packet"

type RequestAuthenticate struct {
	Username string
	Password string
}

func NewRequestAuthenticate(username string, password string) (this *RequestAuthenticate) {
	this = new(RequestAuthenticate)
	this.Username = username
	this.Password = password
	return
}

func (this *RequestAuthenticate) Id() int {
	return REQUEST_AUTHENTICATE
}

type requestAuthenticateCodec struct {

}

func (this *requestAuthenticateCodec) Decode(reader io.Reader) (request Request, err error) {
	requestAuthenticate := new(RequestAuthenticate)
	requestAuthenticate.Username, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	requestAuthenticate.Password, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	request = requestAuthenticate
	return
}

func (this *requestAuthenticateCodec) Encode(writer io.Writer, request Request) (err error) {
	requestAuthenticate := request.(*RequestAuthenticate)
	err = packet.WriteString(writer, requestAuthenticate.Username)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, requestAuthenticate.Password)
	return
}

type ResultAuthenticate struct {

}

func NewResultAuthenticate() (this *ResultAuthenticate) {
	this = new(ResultAuthenticate)
	return
}

func (this *ResultAuthenticate) Id() int {
	return REQUEST_AUTHENTICATE
}

type resultAuthenticateCodec struct {

}

func (this *resultAuthenticateCodec) Decode(reader io.Reader) (result Result, err error) {
	result = new(ResultAuthenticate)
	return
}

func (this *resultAuthenticateCodec) Encode(writer io.Writer, result Result) (err error) {
	return
}
