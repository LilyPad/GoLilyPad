package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientStatusResponse struct {
	Json string
}

func (this *PacketClientStatusResponse) Id() int {
	return PACKET_CLIENT_STATUS_RESPONSE
}

type PacketClientStatusResponseCodec struct {
	
}

func (this *PacketClientStatusResponseCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientStatusResponse := &PacketClientStatusResponse{}
	packetClientStatusResponse.Json, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return packetClientStatusResponse, nil
}

func (this *PacketClientStatusResponseCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientStatusResponse := encode.(*PacketClientStatusResponse)
	err = packet.WriteString(writer, util, packetClientStatusResponse.Json)
	return
}
