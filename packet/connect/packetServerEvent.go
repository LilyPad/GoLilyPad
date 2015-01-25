package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerEvent struct {
	Add bool
	Server string
	SecurityKey string
	Address string
	Port uint16
}

func NewPacketServerEventAdd(server string, securityKey string, address string, port uint16) (this *PacketServerEvent) {
	this = new(PacketServerEvent)
	this.Add = true
	this.Server = server
	this.SecurityKey = securityKey
	this.Address = address
	this.Port = port
	return
}

func NewPacketServerEventRemove(server string) (this *PacketServerEvent) {
	this = new(PacketServerEvent)
	this.Add = false
	this.Server = server
	return
}

func (this *PacketServerEvent) Id() int {
	return PACKET_SERVER_EVENT
}

type packetServerEventCodec struct {

}

func (this *packetServerEventCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerEvent := new(PacketServerEvent)
	packetServerEvent.Add, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetServerEvent.Server, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	if packetServerEvent.Add {
		packetServerEvent.SecurityKey, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetServerEvent.Address, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetServerEvent.Port, err = packet.ReadUint16(reader)
		if err != nil {
			return
		}
	}
	decode = packetServerEvent
	return
}

func (this *packetServerEventCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerEvent := encode.(*PacketServerEvent)
	err = packet.WriteBool(writer, packetServerEvent.Add)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetServerEvent.Server)
	if packetServerEvent.Add {
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetServerEvent.SecurityKey)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetServerEvent.Address)
		if err != nil {
			return
		}
		err = packet.WriteUint16(writer, packetServerEvent.Port)
	}
	return
}
