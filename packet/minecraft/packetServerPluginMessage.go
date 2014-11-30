package minecraft

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	packetServerPluginMessage.Data, err = ioutil.ReadAll(reader)
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
	_, err = writer.Write(packetServerPluginMessage.Data)
	return
}
