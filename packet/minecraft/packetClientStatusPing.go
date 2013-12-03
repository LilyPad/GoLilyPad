package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientStatusPing struct {
	Time int64
}

func (this *PacketClientStatusPing) Id() int {
	return PACKET_CLIENT_STATUS_PING
}

type PacketClientStatusPingCodec struct {
	
}

func (this *PacketClientStatusPingCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientStatusPing := &PacketClientStatusPing{}
	packetClientStatusPing.Time, err = packet.ReadInt64(reader, util)
	if err != nil {
		return
	}
	return packetClientStatusPing, nil
}

func (this *PacketClientStatusPingCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientStatusPing := encode.(*PacketClientStatusPing)
	err = packet.WriteInt64(writer, util, packetClientStatusPing.Time)
	return
}