package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestMessage struct {
	Recipients []string
	Channel string
	Message []byte
}

func (this *RequestMessage) Id() int {
	return REQUEST_MESSAGE
}

type RequestMessageCodec struct {

}

func (this *RequestMessageCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestMessage := &RequestMessage{}
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
	return requestMessage, nil
}

func (this *RequestMessageCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
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

func (this *ResultMessage) Id() int {
	return REQUEST_MESSAGE
}

type ResultMessageCodec struct {

}

func (this *ResultMessageCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	return &ResultMessage{}, nil
}

func (this *ResultMessageCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}
