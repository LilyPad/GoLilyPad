package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketMessageEvent struct {
	Sender string
	Channel string
	Payload []byte
}

func NewPacketMessageEvent(sender string, channel string, payload []byte) (this *PacketMessageEvent) {
	this = new(PacketMessageEvent)
	this.Sender = sender
	this.Channel = channel
	this.Payload = payload
	return
}

func (this *PacketMessageEvent) Id() int {
	return PACKET_MESSAGE_EVENT
}

type packetMessageEventCodec struct {

}

func (this *packetMessageEventCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetMessageEvent := new(PacketMessageEvent)
	packetMessageEvent.Sender, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetMessageEvent.Channel, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	payloadSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetMessageEvent.Payload = make([]byte, payloadSize)
	_, err = reader.Read(packetMessageEvent.Payload)
	if err != nil {
		return
	}
	decode = packetMessageEvent
	return
}

func (this *packetMessageEventCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetMessageEvent := encode.(*PacketMessageEvent)
	err = packet.WriteString(writer, util, packetMessageEvent.Sender)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetMessageEvent.Channel)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetMessageEvent.Payload)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetMessageEvent.Payload)
	return
}
