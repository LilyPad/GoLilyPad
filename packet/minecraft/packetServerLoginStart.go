package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerLoginStart struct {
	Name string
}

func (this *PacketServerLoginStart) Id() int {
	return PACKET_SERVER_LOGIN_START
}

type PacketServerLoginStartCodec struct {
	
}

func (this *PacketServerLoginStartCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerLoginStart := &PacketServerLoginStart{}
	packetServerLoginStart.Name, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return packetServerLoginStart, nil
}

func (this *PacketServerLoginStartCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerLoginStart := encode.(*PacketServerLoginStart)
	err = packet.WriteString(writer, util, packetServerLoginStart.Name)
	return
}
