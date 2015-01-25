package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type packetServerPluginMessageCodec17 struct {

}

func (this *packetServerPluginMessageCodec17) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerPluginMessage := new(PacketServerPluginMessage)
	packetServerPluginMessage.Channel, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	dataLength, err := packet.ReadUint16(reader)
	if err != nil {
		return
	}
	packetServerPluginMessage.Data = make([]byte, dataLength)
	_, err = reader.Read(packetServerPluginMessage.Data)
	if err != nil {
		return
	}
	decode = packetServerPluginMessage
	return
}

func (this *packetServerPluginMessageCodec17) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerPluginMessage := encode.(*PacketServerPluginMessage)
	err = packet.WriteString(writer, packetServerPluginMessage.Channel)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, uint16(len(packetServerPluginMessage.Data)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerPluginMessage.Data)
	return
}
