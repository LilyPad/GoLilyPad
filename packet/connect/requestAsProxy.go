package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type RequestAsProxy struct {
	Address string
	Port uint16
	Motd string
	Version string
	MaxPlayers uint16
}

func NewRequestAsProxy(address string, port uint16, motd string, version string, maxPlayers uint16) (this *RequestAsProxy) {
	this = new(RequestAsProxy)
	this.Address = address
	this.Port = port
	this.Motd = motd
	this.Version = version
	this.MaxPlayers = maxPlayers
	return
}

func (this *RequestAsProxy) Id() int {
	return REQUEST_AS_PROXY
}

type requestAsProxyCodec struct {

}

func (this *requestAsProxyCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestAsProxy := new(RequestAsProxy)
	requestAsProxy.Address, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestAsProxy.Port, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	requestAsProxy.Motd, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestAsProxy.Version, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestAsProxy.MaxPlayers, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	request = requestAsProxy
	return
}

func (this *requestAsProxyCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	requestAsProxy := request.(*RequestAsProxy)
	err = packet.WriteString(writer, util, requestAsProxy.Address)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, requestAsProxy.Port)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, requestAsProxy.Motd)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, requestAsProxy.Version)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, requestAsProxy.MaxPlayers)
	return
}

type ResultAsProxy struct {

}

func NewResultAsProxy() (this *ResultAsProxy) {
	this = new(ResultAsProxy)
	return
}

func (this *ResultAsProxy) Id() int {
	return REQUEST_AS_PROXY
}

type resultAsProxyCodec struct {

}

func (this *resultAsProxyCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	result = new(ResultAsProxy)
	return
}

func (this *resultAsProxyCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}
