package minecraft

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type GameProfile struct {
	Id         string                `json:"id"`
	Name       string                `json:"name"`
	Properties []GameProfileProperty `json:"properties"`
}

type GameProfileProperty struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

func ReadGameProfileProperties(reader io.Reader) (val []GameProfileProperty, err error) {
	var propertiesLength int
	propertiesLength, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if propertiesLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Properties length is below zero: %d", propertiesLength))
		return
	}
	if propertiesLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Properties length is above maximum: %d", propertiesLength))
		return
	}
	properties := make([]GameProfileProperty, propertiesLength)
	for j := range properties {
		property := &properties[j]
		property.Name, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		property.Value, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		var signed bool
		signed, err = packet.ReadBool(reader)
		if err != nil {
			return
		}
		if signed {
			property.Signature, err = packet.ReadString(reader)
			if err != nil {
				return
			}
		}
	}
	val = properties
	return
}

func WriteGameProfileProperties(writer io.Writer, val []GameProfileProperty) (err error) {
	err = packet.WriteVarInt(writer, len(val))
	if err != nil {
		return
	}
	for _, property := range val {
		err = packet.WriteString(writer, property.Name)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, property.Value)
		if err != nil {
			return
		}
		if property.Signature == "" {
			err = packet.WriteBool(writer, false)
			if err != nil {
				return
			}
		} else {
			err = packet.WriteBool(writer, true)
			if err != nil {
				return
			}
			err = packet.WriteString(writer, property.Signature)
			if err != nil {
				return
			}
		}
	}
	return
}
