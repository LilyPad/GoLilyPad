package connect

import "bytes"
import "errors"
import "fmt"
import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketResult struct {
	SequenceId int32
	StatusCode uint8
	Result Result
}

func (this *PacketResult) Id() int {
	return PACKET_RESULT
}

type PacketResultCodec struct {
	sequencer PacketResultSequencer
}

func NewPacketResultCodec(sequencer PacketResultSequencer) *PacketResultCodec {
	return &PacketResultCodec{sequencer}
}

func (this *PacketResultCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	if this.sequencer == nil {
		err = errors.New("No sequencer to decode PacketResult")
		return
	}
	packetResult := &PacketResult{}
	packetResult.SequenceId, err = packet.ReadInt32(reader, util)
	if err != nil {
		return
	}
	packetResult.StatusCode, err = packet.ReadUint8(reader, util)
	if err != nil {
		return
	}
	if packetResult.StatusCode == STATUS_SUCCESS {
		var payloadSize uint16
		payloadSize, err = packet.ReadUint16(reader, util)
		if err != nil {
			return nil, err
		}
		payload := make([]byte, payloadSize)
		_, err = reader.Read(payload)
		if err != nil {
			return
		}
		buffer := bytes.NewBuffer(payload)
		requestId := this.sequencer.RequestIdBySequenceId(packetResult.SequenceId)
		if requestId < 0 {
			err = errors.New(fmt.Sprintf("Request Id is below zero: %i", requestId))
			return
		}
		if int(requestId) >= len(requestCodecs) {
			err = errors.New(fmt.Sprintf("Request Id is above maximum: %i", requestId))
			return
		}
		codec := resultCodecs[requestId]
		if codec == nil {
			err = errors.New(fmt.Sprintf("Request Id does not have a codec: %i", requestId))
			return
		}
		packetResult.Result, err = codec.Decode(buffer, util)
		if err != nil {
			return
		}
	}
	return packetResult, nil
}

func (this *PacketResultCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetResult := encode.(*PacketResult)
	err = packet.WriteInt32(writer, util, packetResult.SequenceId)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, util, packetResult.StatusCode)
	if packetResult.StatusCode == STATUS_SUCCESS {
		if packetResult.Result.Id() < 0 {
			err = errors.New(fmt.Sprintf("Request Id is below zero: %i", packetResult.Result.Id()))
			return
		}
		if packetResult.Result.Id() >= len(requestCodecs) {
			err = errors.New(fmt.Sprintf("Request Id is above maximum: %i", packetResult.Result.Id()))
			return
		}
		buffer := &bytes.Buffer{}
		codec := resultCodecs[packetResult.Result.Id()]
		if codec == nil {
			err = errors.New(fmt.Sprintf("Request Id does not have a codec: %i", packetResult.Result.Id()))
			return
		}
		err = codec.Encode(buffer, util, packetResult.Result)
		if err != nil {
			return
		}
		payload := buffer.Bytes()
		err = packet.WriteUint16(writer, util, uint16(len(payload)))
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
