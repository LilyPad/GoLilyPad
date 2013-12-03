package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerEvent struct {
	Add bool
	Server string
	SecurityKey string
	Address string
	Port uint16
}

func (this *PacketServerEvent) Id() int {
	return PACKET_SERVER_EVENT
}

type PacketServerEventCodec struct {
	
}

func (this *PacketServerEventCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerEvent := &PacketServerEvent{}
	packetServerEvent.Add, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	packetServerEvent.Server, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	if packetServerEvent.Add {
		packetServerEvent.SecurityKey, err = packet.ReadString(reader, util)
		if err != nil {
			return
		}
		packetServerEvent.Address, err = packet.ReadString(reader, util)
		if err != nil {
			return
		}
		packetServerEvent.Port, err = packet.ReadUint16(reader, util)
		if err != nil {
			return
		}
	}
	return packetServerEvent, nil
}

func (this *PacketServerEventCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerEvent := encode.(*PacketServerEvent)
	err = packet.WriteBool(writer, util, packetServerEvent.Add)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetServerEvent.Server)
	if packetServerEvent.Add {
		if err != nil {
			return
		}
		err = packet.WriteString(writer, util, packetServerEvent.SecurityKey)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, util, packetServerEvent.Address)
		if err != nil {
			return
		}
		err = packet.WriteUint16(writer, util, packetServerEvent.Port)
	}
	return
}
