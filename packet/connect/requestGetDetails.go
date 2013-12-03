package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestGetDetails struct {

}

func (this *RequestGetDetails) Id() int {
	return REQUEST_GET_DETAILS
}

type RequestGetDetailsCodec struct {

}

func (this *RequestGetDetailsCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	return &RequestGetDetails{}, nil
}

func (this *RequestGetDetailsCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	return
}

type ResultGetDetails struct {
	Ip string
	Port uint16
	Motd string
	Version string
}

func (this *ResultGetDetails) Id() int {
	return REQUEST_GET_DETAILS
}

type ResultGetDetailsCodec struct {

}

func (this *ResultGetDetailsCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultGetDetails := &ResultGetDetails{}
	resultGetDetails.Ip, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	resultGetDetails.Port, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	resultGetDetails.Motd, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	resultGetDetails.Version, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return resultGetDetails, nil
}

func (this *ResultGetDetailsCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	resultGetDetails := result.(*ResultGetDetails)
	err = packet.WriteString(writer, util, resultGetDetails.Ip)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, resultGetDetails.Port)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, resultGetDetails.Motd)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, resultGetDetails.Version)
	return
}
