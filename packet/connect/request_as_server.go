package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type RequestAsServer struct {
	Address string
	Port    uint16
}

func NewRequestAsServer(address string, port uint16) (this *RequestAsServer) {
	this = new(RequestAsServer)
	this.Address = address
	this.Port = port
	return
}

func (this *RequestAsServer) Id() int {
	return REQUEST_AS_SERVER
}

type requestAsServerCodec struct {
}

func (this *requestAsServerCodec) Decode(reader io.Reader) (request Request, err error) {
	requestAsServer := new(RequestAsServer)
	requestAsServer.Address, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	requestAsServer.Port, err = packet.ReadUint16(reader)
	if err != nil {
		return
	}
	request = requestAsServer
	return
}

func (this *requestAsServerCodec) Encode(writer io.Writer, request Request) (err error) {
	requestAsServer := request.(*RequestAsServer)
	err = packet.WriteString(writer, requestAsServer.Address)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, requestAsServer.Port)
	return
}

type ResultAsServer struct {
	SecurityKey string
}

func NewResultAsServer(securityKey string) (this *ResultAsServer) {
	this = new(ResultAsServer)
	this.SecurityKey = securityKey
	return
}

func (this *ResultAsServer) Id() int {
	return REQUEST_AS_SERVER
}

type resultAsServerCodec struct {
}

func (this *resultAsServerCodec) Decode(reader io.Reader) (result Result, err error) {
	resultAsServer := new(ResultAsServer)
	resultAsServer.SecurityKey, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	result = resultAsServer
	return
}

func (this *resultAsServerCodec) Encode(writer io.Writer, result Result) (err error) {
	err = packet.WriteString(writer, result.(*ResultAsServer).SecurityKey)
	return
}
