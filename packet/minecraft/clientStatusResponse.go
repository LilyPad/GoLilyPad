package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type PacketClientStatusResponse struct {
	Json string
}

func NewPacketClientStatusResponse(json string) (this *PacketClientStatusResponse) {
	this = new(PacketClientStatusResponse)
	this.Json = json
	return
}

func (this *PacketClientStatusResponse) Id() int {
	return PACKET_CLIENT_STATUS_RESPONSE
}

type CodecClientStatusResponse struct {
}

func (this *CodecClientStatusResponse) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientStatusResponse := new(PacketClientStatusResponse)
	packetClientStatusResponse.Json, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetClientStatusResponse
	return
}

func (this *CodecClientStatusResponse) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientStatusResponse := encode.(*PacketClientStatusResponse)
	err = packet.WriteString(writer, packetClientStatusResponse.Json)
	return
}
