package v1193

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	uuid "github.com/satori/go.uuid"
	"io"
)

type CodecClientPlayerInfoUpdate struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientPlayerInfoUpdate) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientPlayerInfoUpdate := new(minecraft.PacketClientPlayerInfoUpdate)
	packetClientPlayerInfoUpdate.IdFrom(this.IdMap)
	bitset, err := packet.ReadFixedBitSet(reader, len(minecraft.PacketClientPlayerInfoUpdateActions))
	if err != nil {
		return
	}
	for _, action := range minecraft.PacketClientPlayerInfoUpdateActions {
		if !bitset.Get(action) {
			continue
		}
		packetClientPlayerInfoUpdate.Actions = append(packetClientPlayerInfoUpdate.Actions, action)
	}
	itemsLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	if itemsLength < 0 {
		err = errors.New(fmt.Sprintf("Decode, Items length is below zero: %d", itemsLength))
		return
	}
	if itemsLength > 65535 {
		err = errors.New(fmt.Sprintf("Decode, Items length is above maximum: %d", itemsLength))
		return
	}
	packetClientPlayerInfoUpdate.Items = make([]minecraft.PacketClientPlayerInfoUpdateItem, itemsLength)
	for i := range packetClientPlayerInfoUpdate.Items {
		item := &packetClientPlayerInfoUpdate.Items[i]
		item.UUID, err = packet.ReadUUID(reader)
		if err != nil {
			return
		}
		for _, action := range packetClientPlayerInfoUpdate.Actions {
			switch action {
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_ADD_PLAYER:
				item.Name, err = packet.ReadString(reader)
				if err != nil {
					return
				}
				item.Properties, err = minecraft.ReadGameProfileProperties(reader)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_INITIALIZE_CHAT:
				var hasChatSession bool
				hasChatSession, err = packet.ReadBool(reader)
				if err != nil {
					return
				}
				if hasChatSession {
					var chatSessionId uuid.UUID
					chatSessionId, err = packet.ReadUUID(reader)
					if err != nil {
						return
					}
					item.ChatSessionId = &chatSessionId
					item.PublicKey, err = minecraft.ReadProfilePublicKey(reader)
					if err != nil {
						return
					}
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_GAMEMODE:
				item.Gamemode, err = packet.ReadVarInt(reader)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_LISTED:
				item.Listed, err = packet.ReadBool(reader)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_LATENCY:
				item.Latency, err = packet.ReadVarInt(reader)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_DISPLAY_NAME:
				var hasDisplayName bool
				hasDisplayName, err = packet.ReadBool(reader)
				if err != nil {
					return
				}
				if hasDisplayName {
					item.DisplayName, err = packet.ReadString(reader)
					if err != nil {
						return
					}
				}
			default:
				err = errors.New(fmt.Sprintf("Decode, PacketClientPlayerInfoUpdateItem action is not valid: %d", action))
				if err != nil {
					return
				}
			}
		}
	}
	decode = packetClientPlayerInfoUpdate
	return
}

func (this *CodecClientPlayerInfoUpdate) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientPlayerInfoUpdate := encode.(*minecraft.PacketClientPlayerInfoUpdate)
	bitset := packet.NewBitSet(len(minecraft.PacketClientPlayerInfoUpdateActions))
	for _, action := range packetClientPlayerInfoUpdate.Actions {
		bitset.Set(action)
	}
	err = packet.WriteFixedBitSet(writer, bitset, len(minecraft.PacketClientPlayerInfoUpdateActions))
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(packetClientPlayerInfoUpdate.Items))
	if err != nil {
		return
	}
	for _, item := range packetClientPlayerInfoUpdate.Items {
		err = packet.WriteUUID(writer, item.UUID)
		if err != nil {
			return
		}
		for _, action := range packetClientPlayerInfoUpdate.Actions {
			switch action {
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_ADD_PLAYER:
				err = packet.WriteString(writer, item.Name)
				if err != nil {
					return
				}
				err = minecraft.WriteGameProfileProperties(writer, item.Properties)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_INITIALIZE_CHAT:
				if item.ChatSessionId == nil {
					err = packet.WriteBool(writer, false)
					if err != nil {
						return
					}
				} else {
					err = packet.WriteBool(writer, true)
					if err != nil {
						return
					}
					err = packet.WriteUUID(writer, *item.ChatSessionId)
					if err != nil {
						return
					}
					err = minecraft.WriteProfilePublicKey(writer, item.PublicKey)
					if err != nil {
						return
					}
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_GAMEMODE:
				err = packet.WriteVarInt(writer, item.Gamemode)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_LISTED:
				err = packet.WriteBool(writer, item.Listed)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_LATENCY:
				err = packet.WriteVarInt(writer, item.Latency)
				if err != nil {
					return
				}
			case minecraft.PACKET_CLIENT_PLAYER_INFO_UPDATE_ACTION_UPDATE_DISPLAY_NAME:
				if item.DisplayName == "" {
					err = packet.WriteBool(writer, false)
					if err != nil {
						return
					}
				} else {
					err = packet.WriteBool(writer, true)
					if err != nil {
						return
					}
					err = packet.WriteString(writer, item.DisplayName)
					if err != nil {
						return
					}
				}
			default:
				err = errors.New(fmt.Sprintf("Encode, PacketClientPlayerInfoUpdateItem action is not valid: %d", action))
				if err != nil {
					return
				}
			}
		}
	}
	return
}
