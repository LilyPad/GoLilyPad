package v17

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type PacketClientPlayerList struct {
	Name   string
	Online bool
	Ping   int16
}

func NewPacketClientPlayerListAdd(name string, ping int16) (this *PacketClientPlayerList) {
	this = new(PacketClientPlayerList)
	this.Name = name
	this.Online = true
	this.Ping = ping
	return
}

func NewPacketClientPlayerListRemove(name string) (this *PacketClientPlayerList) {
	this = new(PacketClientPlayerList)
	this.Name = name
	this.Online = false
	return
}

func (this *PacketClientPlayerList) Id() int {
	return PACKET_CLIENT_PLAYER_LIST
}

type CodecClientPlayerList struct {
}

func (this *CodecClientPlayerList) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientPlayerList := new(PacketClientPlayerList)
	packetClientPlayerList.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientPlayerList.Online, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientPlayerList.Ping, err = packet.ReadInt16(reader)
	if err != nil {
		return
	}
	decode = packetClientPlayerList
	return
}

func (this *CodecClientPlayerList) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientPlayerList := encode.(*PacketClientPlayerList)
	err = packet.WriteString(writer, packetClientPlayerList.Name)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientPlayerList.Online)
	if err != nil {
		return
	}
	err = packet.WriteInt16(writer, packetClientPlayerList.Ping)
	return
}
