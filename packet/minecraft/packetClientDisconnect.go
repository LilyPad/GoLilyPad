package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientDisconnect struct {
	Json string
}

func (this *PacketClientDisconnect) Id() int {
	return PACKET_CLIENT_DISCONNECT
}

type PacketClientDisconnectCodec struct {
	
}

func (this *PacketClientDisconnectCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientDisconnect := &PacketClientDisconnect{}
	packetClientDisconnect.Json, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return packetClientDisconnect, nil
}

func (this *PacketClientDisconnectCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientDisconnect := encode.(*PacketClientDisconnect)
	err = packet.WriteString(writer, util, packetClientDisconnect.Json)
	return
}
