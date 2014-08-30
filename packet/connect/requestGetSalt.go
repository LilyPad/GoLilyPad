package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type RequestGetSalt struct {

}

func NewRequestGetSalt() (this *RequestGetSalt) {
	this = new(RequestGetSalt)
	return
}

func (this *RequestGetSalt) Id() int {
	return REQUEST_GET_SALT
}

type requestGetSaltCodec struct {

}

func (this *requestGetSaltCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	request = new(RequestGetSalt)
	return
}

func (this *requestGetSaltCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	return
}

type ResultGetSalt struct {
	Salt string
}

func NewResultGetSalt(salt string) (this *ResultGetSalt) {
	this = new(ResultGetSalt)
	this.Salt = salt
	return
}

func (this *ResultGetSalt) Id() int {
	return REQUEST_GET_SALT
}

type resultGetSaltCodec struct {

}

func (this *resultGetSaltCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultGetSalt := new(ResultGetSalt)
	resultGetSalt.Salt, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	result = resultGetSalt
	return
}

func (this *resultGetSaltCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	err = packet.WriteString(writer, util, result.(*ResultGetSalt).Salt)
	return
}
