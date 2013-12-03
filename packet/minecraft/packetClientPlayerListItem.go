package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientPlayerListItem struct {
	Name string
	Online bool
	Ping int16
}

func (this *PacketClientPlayerListItem) Id() int {
	return PACKET_CLIENT_PLAYER_LIST_ITEM
}

type PacketClientPlayerListItemCodec struct {
	
}

func (this *PacketClientPlayerListItemCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientPlayerListItem := &PacketClientPlayerListItem{}
	packetClientPlayerListItem.Name, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetClientPlayerListItem.Online, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	packetClientPlayerListItem.Ping, err = packet.ReadInt16(reader, util)
	if err != nil {
		return
	}
	return packetClientPlayerListItem, nil
}

func (this *PacketClientPlayerListItemCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientPlayerListItem := encode.(*PacketClientPlayerListItem)
	err = packet.WriteString(writer, util, packetClientPlayerListItem.Name)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, util, packetClientPlayerListItem.Online)
	if err != nil {
		return
	}
	err = packet.WriteInt16(writer, util, packetClientPlayerListItem.Ping)
	return
}