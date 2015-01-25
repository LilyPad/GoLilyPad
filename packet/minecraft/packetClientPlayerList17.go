package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientPlayerList17 struct {
	Name string
	Online bool
	Ping int16
}

func NewPacketClientPlayerList17Add(name string, ping int16) (this *PacketClientPlayerList17) {
	this = new(PacketClientPlayerList17)
	this.Name = name
	this.Online = true
	this.Ping = ping
	return
}

func NewPacketClientPlayerList17Remove(name string) (this *PacketClientPlayerList17) {
	this = new(PacketClientPlayerList17)
	this.Name = name
	this.Online = false
	return
}

func (this *PacketClientPlayerList17) Id() int {
	return PACKET_CLIENT_PLAYER_LIST
}

type packetClientPlayerListCodec17 struct {

}

func (this *packetClientPlayerListCodec17) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientPlayerList := new(PacketClientPlayerList17)
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

func (this *packetClientPlayerListCodec17) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientPlayerList := encode.(*PacketClientPlayerList17)
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