package v119

import (
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecServerLoginEncryptResponse struct {
	IdMap *minecraft.IdMap
}

func (this *CodecServerLoginEncryptResponse) Decode(reader io.Reader) (packet.Packet, error) {
	result := new(minecraft.PacketServerLoginEncryptResponse)
	result.IdFrom(this.IdMap)

	sharedSecret, err := packet.ReadArrByteLimit(reader, 65535)
	if err != nil {
		return nil, fmt.Errorf("failed to decode shared secret, err: %v", err)
	}
	result.SharedSecret = sharedSecret

	result.DisableSaltAuth, err = packet.ReadBool(reader)
	if err != nil {
		return nil, err
	}
	if !result.DisableSaltAuth {
		result.Salt, err = packet.ReadUint64(reader)
		if err != nil {
			return nil, err
		}
	}

	verifyToken, err := packet.ReadArrByteLimit(reader, 65535)
	if err != nil {
		return nil, fmt.Errorf("failed to deocde verify token, err %v", err)
	}
	result.VerifyToken = verifyToken
	return result, nil
}

func (this *CodecServerLoginEncryptResponse) Encode(writer io.Writer, encode packet.Packet) error {
	writing := encode.(*minecraft.PacketServerLoginEncryptResponse)
	err := packet.WriteArrByte(writer, writing.SharedSecret)
	if err != nil {
		return err
	}
	err = packet.WriteBool(writer, writing.DisableSaltAuth)
	if err != nil {
		return err
	}
	if !writing.DisableSaltAuth {
		err = packet.WriteUint64(writer, writing.Salt)
		if err != nil {
			return err
		}
	}
	err = packet.WriteArrByte(writer, writing.VerifyToken)
	return err
}
