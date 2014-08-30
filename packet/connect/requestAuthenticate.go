package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

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

func (this *requestAuthenticateCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestAuthenticate := new(RequestAuthenticate)
	requestAuthenticate.Username, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestAuthenticate.Password, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	request = requestAuthenticate
	return
}

func (this *requestAuthenticateCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
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

func NewResultAuthenticate() (this *ResultAuthenticate) {
	this = new(ResultAuthenticate)
	return
}

func (this *ResultAuthenticate) Id() int {
	return REQUEST_AUTHENTICATE
}

type resultAuthenticateCodec struct {

}

func (this *resultAuthenticateCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	result = new(ResultAuthenticate)
	return
}

func (this *resultAuthenticateCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}
