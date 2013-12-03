package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientLoginSuccess struct {
	UUID string
	Name string
}

func (this *PacketClientLoginSuccess) Id() int {
	return PACKET_CLIENT_LOGIN_SUCCESS
}

type PacketClientLoginSuccessCodec struct {
	
}

func (this *PacketClientLoginSuccessCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginSuccess := &PacketClientLoginSuccess{}
	packetClientLoginSuccess.UUID, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetClientLoginSuccess.Name, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return packetClientLoginSuccess, nil
}

func (this *PacketClientLoginSuccessCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginSuccess := encode.(*PacketClientLoginSuccess)
	err = packet.WriteString(writer, util, packetClientLoginSuccess.UUID)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetClientLoginSuccess.Name)
	return
}
