package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestGetSalt struct {

}

func (this *RequestGetSalt) Id() int {
	return REQUEST_GET_SALT
}

type RequestGetSaltCodec struct {

}

func (this *RequestGetSaltCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	return &RequestGetSalt{}, nil
}

func (this *RequestGetSaltCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	return
}

type ResultGetSalt struct {
	Salt string
}

func (this *ResultGetSalt) Id() int {
	return REQUEST_GET_SALT
}

type ResultGetSaltCodec struct {
	
}

func (this *ResultGetSaltCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultGetSalt := &ResultGetSalt{}
	resultGetSalt.Salt, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return resultGetSalt, nil
}

func (this *ResultGetSaltCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	err = packet.WriteString(writer, util, result.(*ResultGetSalt).Salt)
	return
}