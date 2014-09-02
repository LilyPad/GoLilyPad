package minecraft

import (
	"errors"
	"fmt"
	"io"
	uuid "code.google.com/p/go-uuid/uuid"
	"github.com/LilyPad/GoLilyPad/packet"
)

const (
	PACKET_CLIENT_PLAYER_LIST_ACTION_ADD = 0
	PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_GAMEMODE = 1
	PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_LATENCY = 2
	PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_DISPLAY_NAME = 3
	PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE = 4
)

type PacketClientPlayerList struct {
	Action int
	Items []PacketClientPlayerListItem
}

type PacketClientPlayerListItem struct {
	UUID uuid.UUID
	Info interface{}
}

type PacketClientPlayerListAddPlayer struct {
	Name string
	Properties []PacketClientPlayerListAddPlayerProperty
	Gamemode int
	Latency int
	DisplayName string
}

type PacketClientPlayerListAddPlayerProperty struct {
	Name string
	Value string
	Signature string
}

type PacketClientPlayerListUpdateGamemode struct {
	Gamemode int
}

type PacketClientPlayerListUpdateLatency struct {
	Latency int
}

type PacketClientPlayerListUpdateDisplayName struct {
	DisplayName string
}

func NewPacketClientPlayerList(action int, items []PacketClientPlayerListItem) (this *PacketClientPlayerList) {
	this = new(PacketClientPlayerList)
	this.Action = action
	this.Items = items
	return
}

func (this *PacketClientPlayerList) Id() int {
	return PACKET_CLIENT_PLAYER_LIST
}

type packetClientPlayerListCodec struct {

}

func (this *packetClientPlayerListCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientPlayerList := new(PacketClientPlayerList)
	packetClientPlayerList.Action, err = packet.ReadVarInt(reader, util)
	if err != nil {
		return
	}
	itemLength, err := packet.ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if itemLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Item length is below zero: %d", itemLength))
		return
	}
	if itemLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Item length is above maximum: %d", itemLength))
		return
	}
	packetClientPlayerList.Items = make([]PacketClientPlayerListItem, itemLength)
	for i, _ := range packetClientPlayerList.Items {
		item := &packetClientPlayerList.Items[i]
		item.UUID, err = packet.ReadUUID(reader, util)
		if err != nil {
			return
		}
		switch packetClientPlayerList.Action {
		case PACKET_CLIENT_PLAYER_LIST_ACTION_ADD:
			addPlayer := PacketClientPlayerListAddPlayer{}
			addPlayer.Name, err = packet.ReadString(reader, util)
			if err != nil {
				return
			}
			var propertiesLength int
			propertiesLength, err = packet.ReadVarInt(reader, util)
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
			addPlayer.Properties = make([]PacketClientPlayerListAddPlayerProperty, propertiesLength)
			for _, property := range addPlayer.Properties {
				property.Name, err = packet.ReadString(reader, util)
				if err != nil {
					return
				}
				property.Value, err = packet.ReadString(reader, util)
				if err != nil {
					return
				}
				var signed bool
				signed, err = packet.ReadBool(reader, util)
				if err != nil {
					return
				}
				if signed {
					property.Signature, err = packet.ReadString(reader, util)
					if err != nil {
						return
					}
				}
			}
			addPlayer.Gamemode, err = packet.ReadVarInt(reader, util)
			if err != nil {
				return
			}
			addPlayer.Latency, err = packet.ReadVarInt(reader, util)
			if err != nil {
				return
			}
			var hasDisplayName bool
			hasDisplayName, err = packet.ReadBool(reader, util)
			if err != nil {
				return
			}
			if hasDisplayName {
				addPlayer.DisplayName, err = packet.ReadString(reader, util)
				if err != nil {
					return
				}
			}
			item.Info = addPlayer
		case PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_GAMEMODE:
			updateGamemode := PacketClientPlayerListUpdateGamemode{}
			updateGamemode.Gamemode, err = packet.ReadVarInt(reader, util)
			if err != nil {
				return
			}
			item.Info = updateGamemode
		case PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_LATENCY:
			updateLatency := PacketClientPlayerListUpdateLatency{}
			updateLatency.Latency, err = packet.ReadVarInt(reader, util)
			if err != nil {
				return
			}
			item.Info = updateLatency
		case PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_DISPLAY_NAME:
			updateDisplayName := PacketClientPlayerListUpdateDisplayName{}
			var hasDisplayName bool
			hasDisplayName, err = packet.ReadBool(reader, util)
			if err != nil {
				return
			}
			if hasDisplayName {
				updateDisplayName.DisplayName, err = packet.ReadString(reader, util)
				if err != nil {
					return
				}
			}
			item.Info = updateDisplayName
		case PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE:
			// no payload
		default:
			err = errors.New(fmt.Sprintf("Decode, PacketClientPlayerList action is not valid: %d", packetClientPlayerList.Action))
		}
	}
	decode = packetClientPlayerList
	return
}

func (this *packetClientPlayerListCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientPlayerList := encode.(*PacketClientPlayerList)
	err = packet.WriteVarInt(writer, util, packetClientPlayerList.Action)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, util, len(packetClientPlayerList.Items))
	if err != nil {
		return
	}
	for _, item := range packetClientPlayerList.Items {
		err = packet.WriteUUID(writer, util, item.UUID)
		if err != nil {
			return
		}
		switch packetClientPlayerList.Action {
		case PACKET_CLIENT_PLAYER_LIST_ACTION_ADD:
			addPlayer := item.Info.(PacketClientPlayerListAddPlayer)
			err = packet.WriteString(writer, util, addPlayer.Name)
			if err != nil {
				return
			}
			err = packet.WriteVarInt(writer, util, len(addPlayer.Properties))
			if err != nil {
				return
			}
			for _, property := range addPlayer.Properties {
				err = packet.WriteString(writer, util, property.Name)
				if err != nil {
					return
				}
				err = packet.WriteString(writer, util, property.Value)
				if err != nil {
					return
				}
				if property.Signature == "" {
					err = packet.WriteBool(writer, util, false)
					if err != nil {
						return
					}
				} else {
					err = packet.WriteBool(writer, util, true)
					if err != nil {
						return
					}
					err = packet.WriteString(writer, util, property.Signature)
					if err != nil {
						return
					}
				}
			}
			err = packet.WriteVarInt(writer, util, addPlayer.Gamemode)
			if err != nil {
				return
			}
			err = packet.WriteVarInt(writer, util, addPlayer.Latency)
			if err != nil {
				return
			}
			if addPlayer.DisplayName == "" {
				err = packet.WriteBool(writer, util, false)
				if err != nil {
					return
				}
			} else {
				err = packet.WriteBool(writer, util, true)
				if err != nil {
					return
				}
				err = packet.WriteString(writer, util, addPlayer.DisplayName)
				if err != nil {
					return
				}
			}
		case PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_GAMEMODE:
			updateGamemode := item.Info.(PacketClientPlayerListUpdateGamemode)
			err = packet.WriteVarInt(writer, util, updateGamemode.Gamemode)
			if err != nil {
				return
			}
		case PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_LATENCY:
			updateLatency := item.Info.(PacketClientPlayerListUpdateLatency)
			err = packet.WriteVarInt(writer, util, updateLatency.Latency)
			if err != nil {
				return
			}
		case PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_DISPLAY_NAME:
			updateDisplayName := item.Info.(PacketClientPlayerListUpdateDisplayName)
			if updateDisplayName.DisplayName == "" {
				err = packet.WriteBool(writer, util, false)
				if err != nil {
					return
				}
			} else {
				err = packet.WriteBool(writer, util, true)
				if err != nil {
					return
				}
				err = packet.WriteString(writer, util, updateDisplayName.DisplayName)
				if err != nil {
					return
				}
			}
		case PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE:
			// no payload
		default:
			err = errors.New(fmt.Sprintf("Encode, PacketClientPlayerList action is not valid: %d", packetClientPlayerList.Action))
		}
	}
	return
}
