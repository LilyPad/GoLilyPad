package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketMessageEvent struct {
	Sender string
	Channel string
	Payload []byte
}

func (this *PacketMessageEvent) Id() int {
	return PACKET_MESSAGE_EVENT
}

type PacketMessageEventCodec struct {
	
}

func (this *PacketMessageEventCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetMessageEvent := &PacketMessageEvent{}
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
	return packetMessageEvent, nil
}

func (this *PacketMessageEventCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
