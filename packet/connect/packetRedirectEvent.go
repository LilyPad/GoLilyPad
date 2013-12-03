package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketRedirectEvent struct {
	Server string
	Player string
}

func (this *PacketRedirectEvent) Id() int {
	return PACKET_REDIRECT_EVENT
}

type PacketRedirectEventCodec struct {
	
}

func (this *PacketRedirectEventCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetRedirectEvent := &PacketRedirectEvent{}
	packetRedirectEvent.Server, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetRedirectEvent.Player, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return packetRedirectEvent, nil
}

func (this *PacketRedirectEventCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetRedirectEvent := encode.(*PacketRedirectEvent)
	err = packet.WriteString(writer, util, packetRedirectEvent.Server)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetRedirectEvent.Player)
	return
}