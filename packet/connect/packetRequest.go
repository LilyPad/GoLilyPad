package connect

import "bytes"
import "errors"
import "fmt"
import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketRequest struct {
	SequenceId int32
	Request Request
}

func (this *PacketRequest) Id() int {
	return PACKET_REQUEST
}

type PacketRequestCodec struct {
	
}

func (this *PacketRequestCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetRequest := &PacketRequest{}
	packetRequest.SequenceId, err = packet.ReadInt32(reader, util)
	if err != nil {
		return
	}
	requestId, err := packet.ReadUint8(reader, util)
	if err != nil {
		return
	}
	if requestId < 0 {
		err = errors.New(fmt.Sprintf("Request Id is below zero: %i", requestId))
		return
	}
	if int(requestId) >= len(requestCodecs) {
		err = errors.New(fmt.Sprintf("Request Id is above maximum: %i", requestId))
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
		err = errors.New(fmt.Sprintf("Request Id does not have a codec: %i", requestId))
		return
	}
	packetRequest.Request, err = codec.Decode(buffer, util)
	if err != nil {
		return
	}
	return packetRequest, nil
}

func (this *PacketRequestCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
		err = errors.New(fmt.Sprintf("Request Id is below zero: %i", packetRequest.Request.Id()))
		return
	}
	if packetRequest.Request.Id() >= len(requestCodecs) {
		err = errors.New(fmt.Sprintf("Request Id is above maximum: %i", packetRequest.Request.Id()))
		return
	}
	buffer := &bytes.Buffer{}
	codec := requestCodecs[packetRequest.Request.Id()]
	if codec == nil {
		err = errors.New(fmt.Sprintf("Request Id does not have a codec: %i", packetRequest.Request.Id()))
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
