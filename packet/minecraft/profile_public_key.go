package minecraft

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type ProfilePublicKey struct {
	Expiry    int64
	Key       []byte
	Signature []byte
}

func ReadProfilePublicKey(reader io.Reader) (val *ProfilePublicKey, err error) {
	expiry, err := packet.ReadInt64(reader)
	if err != nil {
		return
	}
	keyLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if keyLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Key length is below zero: %d", keyLength))
		return
	}
	if keyLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Key length is above maximum: %d", keyLength))
		return
	}
	key := make([]byte, keyLength)
	_, err = reader.Read(key)
	if err != nil {
		return
	}
	signatureLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if signatureLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Signature length is below zero: %d", signatureLength))
		return
	}
	if signatureLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Signature length is above maximum: %d", signatureLength))
		return
	}
	signature := make([]byte, signatureLength)
	_, err = reader.Read(signature)
	if err != nil {
		return
	}
	val = &ProfilePublicKey{
		Expiry:    expiry,
		Key:       key,
		Signature: signature,
	}
	return
}

func WriteProfilePublicKey(writer io.Writer, val *ProfilePublicKey) (err error) {
	err = packet.WriteInt64(writer, val.Expiry)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(val.Key))
	if err != nil {
		return
	}
	_, err = writer.Write(val.Key)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(val.Signature))
	if err != nil {
		return
	}
	_, err = writer.Write(val.Signature)
	return
}
