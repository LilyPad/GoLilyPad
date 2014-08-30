package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type RequestMessage struct {
	Recipients []string
	Channel string
	Message []byte
}

func NewRequestMessage(recipients []string, channel string, message []byte) (this *RequestMessage) {
	this = new(RequestMessage)
	this.Recipients = recipients
	this.Channel = channel
	this.Message = message
	return
}

func (this *RequestMessage) Id() int {
	return REQUEST_MESSAGE
}

type requestMessageCodec struct {

}

func (this *requestMessageCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestMessage := new(RequestMessage)
	recipientsSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	requestMessage.Recipients = make([]string, recipientsSize)
	var i uint16
	for i = 0; i < recipientsSize; i++ {
		requestMessage.Recipients[i], err = packet.ReadString(reader, util)
		if err != nil {
			return
		}
	}
	requestMessage.Channel, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	messageSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	requestMessage.Message = make([]byte, messageSize)
	_, err = reader.Read(requestMessage.Message)
	if err != nil {
		return
	}
	request = requestMessage
	return
}

func (this *requestMessageCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	requestMessage := request.(*RequestMessage)
	err = packet.WriteUint16(writer, util, uint16(len(requestMessage.Recipients)))
	for i := 0; i < len(requestMessage.Recipients); i++ {
		if err != nil {
			return
		}
		err = packet.WriteString(writer, util, requestMessage.Recipients[i])
	}
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, requestMessage.Channel)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(requestMessage.Message)))
	if err != nil {
		return
	}
	_, err = writer.Write(requestMessage.Message)
	return
}

type ResultMessage struct {

}

func NewResultMessage() (this *ResultMessage) {
	this = new(ResultMessage)
	return
}

func (this *ResultMessage) Id() int {
	return REQUEST_MESSAGE
}

type resultMessageCodec struct {

}

func (this *resultMessageCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	result = new(ResultMessage)
	return
}

func (this *resultMessageCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}
