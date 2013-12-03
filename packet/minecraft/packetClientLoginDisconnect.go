package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientLoginDisconnect struct {
	Json string
}

func (this *PacketClientLoginDisconnect) Id() int {
	return PACKET_CLIENT_LOGIN_DISCONNECT
}

type PacketClientLoginDisconnectCodec struct {
	
}

func (this *PacketClientLoginDisconnectCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginDisconnect := &PacketClientLoginDisconnect{}
	packetClientLoginDisconnect.Json, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return packetClientLoginDisconnect, nil
}

func (this *PacketClientLoginDisconnectCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginDisconnect := encode.(*PacketClientLoginDisconnect)
	err = packet.WriteString(writer, util, packetClientLoginDisconnect.Json)
	return
}
