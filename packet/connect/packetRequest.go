package connect

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketRequest struct {
	SequenceId int32
	Request Request
}

func NewPacketRequest(sequenceId int32, request Request) (this *PacketRequest) {
	this = new(PacketRequest)
	this.SequenceId = sequenceId
	this.Request = request
	return
}

func (this *PacketRequest) Id() int {
	return PACKET_REQUEST
}

type packetRequestCodec struct {

}

func (this *packetRequestCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetRequest := new(PacketRequest)
	packetRequest.SequenceId, err = packet.ReadInt32(reader, util)
	if err != nil {
		return
	}
	requestId, err := packet.ReadUint8(reader, util)
	if err != nil {
		return
	}
	if requestId < 0 {
		err = errors.New(fmt.Sprintf("Decode, Request Id is below zero: %d", requestId))
		return
	}
	if int(requestId) >= len(requestCodecs) {
		err = errors.New(fmt.Sprintf("Decode, Request Id is above maximum: %d", requestId))
		return
	}
	payloadSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	payload := make([]byte, payloadSize)
	_, err = reader.Read(payload)
	if err != nil {
		return
	}
	buffer := bytes.NewBuffer(payload)
	codec := requestCodecs[requestId]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Decode, Request Id does not have a codec: %d", requestId))
		return
	}
	packetRequest.Request, err = codec.Decode(buffer, util)
	if err != nil {
		return
	}
	decode = packetRequest
	return
}

func (this *packetRequestCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetRequest := encode.(*PacketRequest)
	err = packet.WriteInt32(writer, util, packetRequest.SequenceId)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, util, uint8(packetRequest.Request.Id()))
	if err != nil {
		return
	}
	if packetRequest.Request.Id() < 0 {
		err = errors.New(fmt.Sprintf("Encode, Request Id is below zero: %d", packetRequest.Request.Id()))
		return
	}
	if packetRequest.Request.Id() >= len(requestCodecs) {
		err = errors.New(fmt.Sprintf("Encode, Request Id is above maximum: %d", packetRequest.Request.Id()))
		return
	}
	buffer := new(bytes.Buffer)
	codec := requestCodecs[packetRequest.Request.Id()]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Encode, Request Id does not have a codec: %d", packetRequest.Request.Id()))
		return
	}
	err = codec.Encode(buffer, util, packetRequest.Request)
	if err != nil {
		return
	}
	payload := buffer.Bytes()
	err = packet.WriteUint16(writer, util, uint16(len(payload)))
	if err != nil {
		return
	}
	_, err = writer.Write(payload)
	return
}
