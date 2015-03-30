package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
	"strings"
)

type PacketServerHandshake struct {
	ProtocolVersion int
	ServerAddress string
	ServerPort uint16
	State int
	NullAppendedData string
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

func (this *packetServerHandshakeCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerHandshake := new(PacketServerHandshake)
	packetServerHandshake.ProtocolVersion, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetServerHandshake.ServerAddress, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetServerHandshake.ServerPort, err = packet.ReadUint16(reader)
	if err != nil {
		return
	}
	packetServerHandshake.State, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}

	// Handling for 1.8 FML appending '\0FML\0' to the host-string
	idx := strings.Index(packetServerHandshake.ServerAddress, "\x00")
	if idx != -1 {
		packetServerHandshake.NullAppendedData = packetServerHandshake.ServerAddress[idx:]
		packetServerHandshake.ServerAddress = packetServerHandshake.ServerAddress[:idx]
	}

	decode = packetServerHandshake
	return
}

func (this *packetServerHandshakeCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerHandshake := encode.(*PacketServerHandshake)
	err = packet.WriteVarInt(writer, packetServerHandshake.ProtocolVersion)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetServerHandshake.ServerAddress)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, packetServerHandshake.ServerPort)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, packetServerHandshake.State)
	return
}
