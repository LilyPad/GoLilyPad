package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerLoginEncryptResponse struct {
	SharedSecret []byte
	VerifyToken []byte
}

func (this *PacketServerLoginEncryptResponse) Id() int {
	return PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE
}

type PacketServerLoginEncryptResponseCodec struct {
	
}

func (this *PacketServerLoginEncryptResponseCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerLoginEncryptResponse := &PacketServerLoginEncryptResponse{}
	sharedSecretSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.SharedSecret = make([]byte, sharedSecretSize)
	_, err = reader.Read(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	verifyTokenSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetServerLoginEncryptResponse.VerifyToken = make([]byte, verifyTokenSize)
	_, err = reader.Read(packetServerLoginEncryptResponse.VerifyToken)
	if err != nil {
		return
	}
	return packetServerLoginEncryptResponse, nil
}

func (this *PacketServerLoginEncryptResponseCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerLoginEncryptResponse := encode.(*PacketServerLoginEncryptResponse)
	err = packet.WriteUint16(writer, util, uint16(len(packetServerLoginEncryptResponse.SharedSecret)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.SharedSecret)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetServerLoginEncryptResponse.VerifyToken)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetServerLoginEncryptResponse.VerifyToken)
	return
}