package minecraft

import (
	"errors"
	"fmt"
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerPluginMessage struct {
	Channel string
	Data []byte
}

func NewPacketServerPluginMessage(channel string, data []byte) (this *PacketServerPluginMessage) {
	this = new(PacketServerPluginMessage)
	this.Channel = channel
	this.Data = data
	return
}

func (this *PacketServerPluginMessage) Id() int {
	return PACKET_SERVER_PLUGIN_MESSAGE
}

type packetServerPluginMessageCodec struct {

}

func (this *packetServerPluginMessageCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerPluginMessage := new(PacketServerPluginMessage)
	packetServerPluginMessage.Channel, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	dataLength, err := packet.ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if dataLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Data length is below zero: %d", dataLength))
		return
	}
	if dataLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Data length is above maximum: %d", dataLength))
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

func (this *packetServerPluginMessageCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerPluginMessage := encode.(*PacketServerPluginMessage)
	err = packet.WriteString(writer, util, packetServerPluginMessage.Channel)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, util, len(packetServerPluginMessage.Data))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerPluginMessage.Data)
	return
}
