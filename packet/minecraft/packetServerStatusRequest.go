package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerStatusRequest struct {

}

func (this *PacketServerStatusRequest) Id() int {
	return PACKET_SERVER_STATUS_REQUEST
}

type PacketServerStatusRequestCodec struct {
	
}

func (this *PacketServerStatusRequestCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	return &PacketServerStatusRequest{}, nil
}

func (this *PacketServerStatusRequestCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	return
}