package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerPluginMessage struct {
	Channel string
	Data []byte
}

func (this *PacketServerPluginMessage) Id() int {
	return PACKET_SERVER_PLUGIN_MESSAGE
}

type PacketServerPluginMessageCodec struct {
	
}

func (this *PacketServerPluginMessageCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerPluginMessage := &PacketServerPluginMessage{}
	packetServerPluginMessage.Channel, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	dataSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerPluginMessage.Data = make([]byte, dataSize)
	_, err = reader.Read(packetServerPluginMessage.Data)
	if err != nil {
		return
	}
	return packetServerPluginMessage, nil
}

func (this *PacketServerPluginMessageCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerPluginMessage := encode.(*PacketServerPluginMessage)
	err = packet.WriteString(writer, util, packetServerPluginMessage.Channel)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetServerPluginMessage.Data)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerPluginMessage.Data)
	return
}
