package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestAsServer struct {
	Address string
	Port uint16
}

func (this *RequestAsServer) Id() int {
	return REQUEST_AS_SERVER
}

type RequestAsServerCodec struct {

}

func (this *RequestAsServerCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestAsServer := &RequestAsServer{}
	requestAsServer.Address, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	requestAsServer.Port, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	return requestAsServer, nil
}

func (this *RequestAsServerCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	requestAsServer := request.(*RequestAsServer)
	err = packet.WriteString(writer, util, requestAsServer.Address)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, requestAsServer.Port)
	return
}

type ResultAsServer struct {
	SecurityKey string
}

func (this *ResultAsServer) Id() int {
	return REQUEST_AS_SERVER
}

type ResultAsServerCodec struct {

}

func (this *ResultAsServerCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultAsServer := &ResultAsServer{}
	resultAsServer.SecurityKey, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return resultAsServer, nil
}

func (this *ResultAsServerCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	err = packet.WriteString(writer, util, result.(*ResultAsServer).SecurityKey)
	return
}
