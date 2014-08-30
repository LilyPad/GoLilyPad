package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientPlayerListItem struct {
	Name string
	Online bool
	Ping int16
}

func NewPacketClientPlayerListItemAdd(name string, ping int16) (this *PacketClientPlayerListItem) {
	this = new(PacketClientPlayerListItem)
	this.Name = name
	this.Online = true
	this.Ping = ping
	return
}

func NewPacketClientPlayerListItemRemove(name string) (this *PacketClientPlayerListItem) {
	this = new(PacketClientPlayerListItem)
	this.Name = name
	this.Online = false
	return
}

func (this *PacketClientPlayerListItem) Id() int {
	return PACKET_CLIENT_PLAYER_LIST_ITEM
}

type packetClientPlayerListItemCodec struct {

}

func (this *packetClientPlayerListItemCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientPlayerListItem := new(PacketClientPlayerListItem)
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
	decode = packetClientPlayerListItem
	return
}

func (this *packetClientPlayerListItemCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
