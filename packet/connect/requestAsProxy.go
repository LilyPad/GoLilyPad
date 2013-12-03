package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestAsProxy struct {
	Address string
	Port uint16
	Motd string
	Version string
	Maxplayers uint16
}

func (this *RequestAsProxy) Id() int {
	return REQUEST_AS_PROXY
}

type RequestAsProxyCodec struct {

}

func (this *RequestAsProxyCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestAsProxy := &RequestAsProxy{}
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
	requestAsProxy.Maxplayers, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	return requestAsProxy, nil
}

func (this *RequestAsProxyCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
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
	err = packet.WriteUint16(writer, util, requestAsProxy.Maxplayers)
	return
}

type ResultAsProxy struct {

}

func (this *ResultAsProxy) Id() int {
	return REQUEST_AS_PROXY
}

type ResultAsProxyCodec struct {

}

func (this *ResultAsProxyCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	return &ResultAsProxy{}, nil
}

func (this *ResultAsProxyCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}