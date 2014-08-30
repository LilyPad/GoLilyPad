package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerHandshake struct {
	ProtocolVersion int
	ServerAddress string
	ServerPort uint16
	State int
}

func NewPacketServerHandshake(protocolVersion int, serverAddress string, serverPort uint16, state int) (this *PacketServerHandshake) {
	this = new(PacketServerHandshake)
	this.ProtocolVersion = protocolVersion
	this.ServerAddress = serverAddress
	this.ServerPort = serverPort
	this.State = state
	return
}

func (this *PacketServerHandshake) Id() int {
	return PACKET_SERVER_HANDSHAKE
}

type packetServerHandshakeCodec struct {

}

func (this *packetServerHandshakeCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerHandshake := new(PacketServerHandshake)
	packetServerHandshake.ProtocolVersion, err = packet.ReadVarInt(reader, util)
	if err != nil {
		return
	}
	packetServerHandshake.ServerAddress, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetServerHandshake.ServerPort, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerHandshake.State, err = packet.ReadVarInt(reader, util)
	if err != nil {
		return
	}
	decode = packetServerHandshake
	return
}

func (this *packetServerHandshakeCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerHandshake := encode.(*PacketServerHandshake)
	err = packet.WriteVarInt(writer, util, packetServerHandshake.ProtocolVersion)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetServerHandshake.ServerAddress)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, packetServerHandshake.ServerPort)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, util, packetServerHandshake.State)
	return
}
