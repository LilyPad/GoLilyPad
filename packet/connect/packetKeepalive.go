package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketKeepalive struct {
	Random int32
}

func (this *PacketKeepalive) Id() int {
	return PACKET_KEEPALIVE
}

type PacketKeepaliveCodec struct {
	
}

func (this *PacketKeepaliveCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetKeepalive := &PacketKeepalive{}
	packetKeepalive.Random, err = packet.ReadInt32(reader, util)
	if err != nil {
		return
	}
	return packetKeepalive, nil
}

func (this *PacketKeepaliveCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	err = packet.WriteInt32(writer, util, encode.(*PacketKeepalive).Random)
	return
}
