package v119

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientPlayerList struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientPlayerList) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientPlayerList := new(minecraft.PacketClientPlayerList)
	packetClientPlayerList.IdFrom(this.IdMap)
	packetClientPlayerList.Action, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	itemLength, err := packet.ReadVarInt(reader)
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
	packetClientPlayerList.Items = make([]minecraft.PacketClientPlayerListItem, itemLength)
	for i := range packetClientPlayerList.Items {
		item := &packetClientPlayerList.Items[i]
		item.UUID, err = packet.ReadUUID(reader)
		if err != nil {
			return
		}
		switch packetClientPlayerList.Action {
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_ADD:
			addPlayer := minecraft.PacketClientPlayerListAddPlayer{}
			addPlayer.Name, err = packet.ReadString(reader)
			if err != nil {
				return
			}
			addPlayer.Properties, err = minecraft.ReadGameProfileProperties(reader)
			if err != nil {
				return
			}
			addPlayer.Gamemode, err = packet.ReadVarInt(reader)
			if err != nil {
				return
			}
			addPlayer.Latency, err = packet.ReadVarInt(reader)
			if err != nil {
				return
			}
			var hasDisplayName bool
			hasDisplayName, err = packet.ReadBool(reader)
			if err != nil {
				return
			}
			if hasDisplayName {
				addPlayer.DisplayName, err = packet.ReadString(reader)
				if err != nil {
					return
				}
			}
			var publicKeyPresent bool
			publicKeyPresent, err = packet.ReadBool(reader)
			if err != nil {
				return
			}
			if publicKeyPresent {
				addPlayer.PublicKey, err = minecraft.ReadProfilePublicKey(reader)
				if err != nil {
					return
				}
			}
			item.Info = addPlayer
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_GAMEMODE:
			updateGamemode := minecraft.PacketClientPlayerListUpdateGamemode{}
			updateGamemode.Gamemode, err = packet.ReadVarInt(reader)
			if err != nil {
				return
			}
			item.Info = updateGamemode
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_LATENCY:
			updateLatency := minecraft.PacketClientPlayerListUpdateLatency{}
			updateLatency.Latency, err = packet.ReadVarInt(reader)
			if err != nil {
				return
			}
			item.Info = updateLatency
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_DISPLAY_NAME:
			updateDisplayName := minecraft.PacketClientPlayerListUpdateDisplayName{}
			var hasDisplayName bool
			hasDisplayName, err = packet.ReadBool(reader)
			if err != nil {
				return
			}
			if hasDisplayName {
				updateDisplayName.DisplayName, err = packet.ReadString(reader)
				if err != nil {
					return
				}
			}
			item.Info = updateDisplayName
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE:
			// no payload
		default:
			err = errors.New(fmt.Sprintf("Decode, PacketClientPlayerList action is not valid: %d", packetClientPlayerList.Action))
		}
	}
	decode = packetClientPlayerList
	return
}

func (this *CodecClientPlayerList) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientPlayerList := encode.(*minecraft.PacketClientPlayerList)
	err = packet.WriteVarInt(writer, packetClientPlayerList.Action)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(packetClientPlayerList.Items))
	if err != nil {
		return
	}
	for _, item := range packetClientPlayerList.Items {
		err = packet.WriteUUID(writer, item.UUID)
		if err != nil {
			return
		}
		switch packetClientPlayerList.Action {
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_ADD:
			addPlayer := item.Info.(minecraft.PacketClientPlayerListAddPlayer)
			err = packet.WriteString(writer, addPlayer.Name)
			if err != nil {
				return
			}
			err = minecraft.WriteGameProfileProperties(writer, addPlayer.Properties)
			if err != nil {
				return
			}
			err = packet.WriteVarInt(writer, addPlayer.Gamemode)
			if err != nil {
				return
			}
			err = packet.WriteVarInt(writer, addPlayer.Latency)
			if err != nil {
				return
			}
			if addPlayer.DisplayName == "" {
				err = packet.WriteBool(writer, false)
				if err != nil {
					return
				}
			} else {
				err = packet.WriteBool(writer, true)
				if err != nil {
					return
				}
				err = packet.WriteString(writer, addPlayer.DisplayName)
				if err != nil {
					return
				}
			}
			if addPlayer.PublicKey == nil {
				err = packet.WriteBool(writer, false)
				if err != nil {
					return
				}
			} else {
				err = packet.WriteBool(writer, true)
				if err != nil {
					return
				}
				err = minecraft.WriteProfilePublicKey(writer, addPlayer.PublicKey)
				if err != nil {
					return
				}
			}
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_GAMEMODE:
			updateGamemode := item.Info.(minecraft.PacketClientPlayerListUpdateGamemode)
			err = packet.WriteVarInt(writer, updateGamemode.Gamemode)
			if err != nil {
				return
			}
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_LATENCY:
			updateLatency := item.Info.(minecraft.PacketClientPlayerListUpdateLatency)
			err = packet.WriteVarInt(writer, updateLatency.Latency)
			if err != nil {
				return
			}
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_DISPLAY_NAME:
			updateDisplayName := item.Info.(minecraft.PacketClientPlayerListUpdateDisplayName)
			if updateDisplayName.DisplayName == "" {
				err = packet.WriteBool(writer, false)
				if err != nil {
					return
				}
			} else {
				err = packet.WriteBool(writer, true)
				if err != nil {
					return
				}
				err = packet.WriteString(writer, updateDisplayName.DisplayName)
				if err != nil {
					return
				}
			}
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE:
			// no payload
		default:
			err = errors.New(fmt.Sprintf("Encode, PacketClientPlayerList action is not valid: %d", packetClientPlayerList.Action))
		}
	}
	return
}
