package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerStatusPing struct {
	Time int64
}

func (this *PacketServerStatusPing) Id() int {
	return PACKET_SERVER_STATUS_PING
}

type PacketServerStatusPingCodec struct {
	
}

func (this *PacketServerStatusPingCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerStatusPing := &PacketServerStatusPing{}
	packetServerStatusPing.Time, err = packet.ReadInt64(reader, util)
	if err != nil {
		return
	}
	return packetServerStatusPing, nil
}

func (this *PacketServerStatusPingCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerStatusPing := encode.(*PacketServerStatusPing)
	err = packet.WriteInt64(writer, util, packetServerStatusPing.Time)
	return
}