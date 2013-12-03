package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientLoginEncryptRequest struct {
	ServerId string
	PublicKey []byte
	VerifyToken []byte
}

func (this *PacketClientLoginEncryptRequest) Id() int {
	return PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST
}

type PacketClientLoginEncryptRequestCodec struct {
	
}

func (this *PacketClientLoginEncryptRequestCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientLoginEncryptRequest := &PacketClientLoginEncryptRequest{}
	packetClientLoginEncryptRequest.ServerId, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	publicKeySize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetClientLoginEncryptRequest.PublicKey = make([]byte, publicKeySize)
	_, err = reader.Read(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	verifyTokenSize, err := packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	packetClientLoginEncryptRequest.VerifyToken = make([]byte, verifyTokenSize)
	_, err = reader.Read(packetClientLoginEncryptRequest.VerifyToken)
	if err != nil {
		return
	}
	return packetClientLoginEncryptRequest, nil
}

func (this *PacketClientLoginEncryptRequestCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientLoginEncryptRequest := encode.(*PacketClientLoginEncryptRequest)
	err = packet.WriteString(writer, util, packetClientLoginEncryptRequest.ServerId)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetClientLoginEncryptRequest.PublicKey)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.PublicKey)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, uint16(len(packetClientLoginEncryptRequest.VerifyToken)))
	if err != nil {
		return
	}
	_, err = writer.Write(packetClientLoginEncryptRequest.VerifyToken)
	return
}