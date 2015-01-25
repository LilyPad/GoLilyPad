package connect

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketResult struct {
	SequenceId int32
	StatusCode uint8
	Result Result
}

func NewPacketResult(sequenceId int32, statusCode uint8, result Result) (this *PacketResult) {
	this = new(PacketResult)
	this.SequenceId = sequenceId
	this.StatusCode = statusCode
	this.Result = result
	return
}

func (this *PacketResult) Id() int {
	return PACKET_RESULT
}

type packetResultCodec struct {
	Sequencer PacketResultSequencer
}

func NewPacketResultCodec(sequencer PacketResultSequencer) (this *packetResultCodec) {
	this = new(packetResultCodec)
	this.Sequencer = sequencer
	return
}

func (this *packetResultCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	if this.Sequencer == nil {
		err = errors.New("No sequencer to decode PacketResult")
		return
	}
	packetResult := new(PacketResult)
	packetResult.SequenceId, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetResult.StatusCode, err = packet.ReadUint8(reader)
	if err != nil {
		return
	}
	if packetResult.StatusCode == STATUS_SUCCESS {
		var payloadSize uint16
		payloadSize, err = packet.ReadUint16(reader)
		if err != nil {
			return
		}
		payload := make([]byte, payloadSize)
		_, err = reader.Read(payload)
		if err != nil {
			return
		}
		buffer := bytes.NewBuffer(payload)
		requestId := this.Sequencer.RequestIdBySequenceId(packetResult.SequenceId)
		if requestId < 0 {
			err = errors.New(fmt.Sprintf("Decode, Request Id is below zero: %d", requestId))
			return
		}
		if int(requestId) >= len(requestCodecs) {
			err = errors.New(fmt.Sprintf("Decode, Request Id is above maximum: %d", requestId))
			return
		}
		codec := resultCodecs[requestId]
		if codec == nil {
			err = errors.New(fmt.Sprintf("Decode, Request Id does not have a codec: %d", requestId))
			return
		}
		packetResult.Result, err = codec.Decode(buffer)
		if err != nil {
			return
		}
	}
	decode = packetResult
	return
}

func (this *packetResultCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetResult := encode.(*PacketResult)
	err = packet.WriteInt32(writer, packetResult.SequenceId)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, packetResult.StatusCode)
	if packetResult.StatusCode == STATUS_SUCCESS {
		if packetResult.Result.Id() < 0 {
			err = errors.New(fmt.Sprintf("Encode, Request Id is below zero: %d", packetResult.Result.Id()))
			return
		}
		if packetResult.Result.Id() >= len(requestCodecs) {
			err = errors.New(fmt.Sprintf("Encode, Request Id is above maximum: %d", packetResult.Result.Id()))
			return
		}
		buffer := new(bytes.Buffer)
		codec := resultCodecs[packetResult.Result.Id()]
		if codec == nil {
			err = errors.New(fmt.Sprintf("Encode, Request Id does not have a codec: %d", packetResult.Result.Id()))
			return
		}
		err = codec.Encode(buffer, packetResult.Result)
		if err != nil {
			return
		}
		payload := buffer.Bytes()
		err = packet.WriteUint16(writer, uint16(len(payload)))
		if err != nil {
			return
		}
		_, err = writer.Write(payload)
	}
	return
}

type PacketResultSequencer interface {
	RequestIdBySequenceId(int32) int
}
