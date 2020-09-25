package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
	"io/ioutil"
)

type CodecServerPluginMessage struct {
	IdMap *minecraft.IdMap
}

func (this *CodecServerPluginMessage) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerPluginMessage := new(minecraft.PacketServerPluginMessage)
	packetServerPluginMessage.IdFrom(this.IdMap)
	packetServerPluginMessage.Channel, err = packet.ReadString(reader)
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

func (this *CodecServerPluginMessage) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerPluginMessage := encode.(*minecraft.PacketServerPluginMessage)
	err = packet.WriteString(writer, packetServerPluginMessage.Channel)
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerPluginMessage.Data)
	return
}
