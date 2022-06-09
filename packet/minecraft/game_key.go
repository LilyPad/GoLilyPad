package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type GameKey struct {
	Expiry    uint64
	Key       []byte
	Signature []byte
}

func ReadGameKey(reader io.Reader) (*GameKey, error) {
	expiry, err := packet.ReadUint64(reader)
	if err != nil {
		return nil, err
	}
	key, err := packet.ReadArrByte(reader)
	if err != nil {
		return nil, err
	}
	signature, err := packet.ReadArrByteLimit(reader, 4096)
	if err != nil {
		return nil, err
	}
	return &GameKey{Expiry: expiry, Key: key, Signature: signature}, nil
}

func WriteGameKey(writer io.Writer, key *GameKey) error {
	err := packet.WriteUint64(writer, key.Expiry)
	if err != nil {
		return err
	}
	err = packet.WriteArrByte(writer, key.Key)
	if err != nil {
		return err
	}
	return packet.WriteArrByte(writer, key.Signature)
}
