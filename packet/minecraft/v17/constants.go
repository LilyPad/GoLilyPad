package v17

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	mc18 "github.com/LilyPad/GoLilyPad/packet/minecraft/v18"
)

const (
	PACKET_CLIENT_KEEPALIVE                     = 0x00
	PACKET_CLIENT_JOIN_GAME                     = 0x01
	PACKET_CLIENT_CHAT                          = 0x02
	PACKET_CLIENT_TIME_UPDATE                   = 0x03
	PACKET_CLIENT_ENTITY_EQUIPMENT              = 0x04
	PACKET_CLIENT_SPAWN_POSITION                = 0x05
	PACKET_CLIENT_UPDATE_HEALTH                 = 0x06
	PACKET_CLIENT_RESPAWN                       = 0x07
	PACKET_CLIENT_PLAYER_POSITION_AND_LOOK      = 0x08
	PACKET_CLIENT_HELD_ITEM_CHANGE              = 0x09
	PACKET_CLIENT_USE_BED                       = 0x0A
	PACKET_CLIENT_ANIMATION                     = 0x0B
	PACKET_CLIENT_SPAWN_PLAYER                  = 0x0C
	PACKET_CLIENT_COLLECT_ITEM                  = 0x0D
	PACKET_CLIENT_SPAWN_OBJECT                  = 0x0E
	PACKET_CLIENT_SPAWN_MOB                     = 0x0F
	PACKET_CLIENT_SPAWN_PAINTING                = 0x10
	PACKET_CLIENT_SPAWN_EXPERIENCE_ORB          = 0x11
	PACKET_CLIENT_ENTITY_VELOCITY               = 0x12
	PACKET_CLIENT_DESTROY_ENTITIES              = 0x13
	PACKET_CLIENT_ENTITY                        = 0x14
	PACKET_CLIENT_ENTITY_RELATIVE_MOVE          = 0x15
	PACKET_CLIENT_ENTITY_LOOK                   = 0x16
	PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE = 0x17
	PACKET_CLIENT_ENTITY_TELEPORT               = 0x18
	PACKET_CLIENT_ENTITY_HEAD_LOOK              = 0x19
	PACKET_CLIENT_ENTITY_STATUS                 = 0x1A
	PACKET_CLIENT_ATTACH_ENTITY                 = 0x1B
	PACKET_CLIENT_ENTITY_METADATA               = 0x1C
	PACKET_CLIENT_ENTITY_EFFECT                 = 0x1D
	PACKET_CLIENT_REMOVE_ENTITY_EFFECT          = 0x1E
	PACKET_CLIENT_SET_EXPERIENCE                = 0x1F
	PACKET_CLIENT_ENTITY_PROPERTIES             = 0x20
	PACKET_CLIENT_CHUNK_DATA                    = 0x21
	PACKET_CLIENT_MULTI_BLOCK_CHANGE            = 0x22
	PACKET_CLIENT_BLOCK_CHANGE                  = 0x23
	PACKET_CLIENT_BLOCK_ACTION                  = 0x24
	PACKET_CLIENT_BLOCK_BREAK_ANIMATION         = 0x25
	PACKET_CLIENT_MAP_CHUNK_BULK                = 0x26
	PACKET_CLIENT_EXPLOSION                     = 0x27
	PACKET_CLIENT_EFFECT                        = 0x28
	PACKET_CLIENT_NAMED_SOUND_EFFECT            = 0x29
	PACKET_CLIENT_PARTICLE                      = 0x2A
	PACKET_CLIENT_CHANGE_GAME_STATE             = 0x2B
	PACKET_CLIENT_SPAWN_GLOBAL_ENTITY           = 0x2C
	PACKET_CLIENT_OPEN_WINDOW                   = 0x2D
	PACKET_CLIENT_CLOSE_WINDOW                  = 0x2E
	PACKET_CLIENT_SET_SLOT                      = 0x2F
	PACKET_CLIENT_WINDOW_ITEMS                  = 0x30
	PACKET_CLIENT_WINDOW_PROPERTY               = 0x31
	PACKET_CLIENT_CONFIRM_TRANSACTION           = 0x32
	PACKET_CLIENT_UPDATE_SIGN                   = 0x33
	PACKET_CLIENT_MAPS                          = 0x34
	PACKET_CLIENT_UPDATE_BLOCK_ENTITY           = 0x35
	PACKET_CLIENT_SIGN_EDITOR_OPEN              = 0x36
	PACKET_CLIENT_STATISTICS                    = 0x37
	PACKET_CLIENT_PLAYER_LIST                   = 0x38
	PACKET_CLIENT_PLAYER_ABILITIES              = 0x39
	PACKET_CLIENT_TAB_COMPLETE                  = 0x3A
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE          = 0x3B
	PACKET_CLIENT_UPDATE_SCORE                  = 0x3C
	PACKET_CLIENT_DISPLAY_SCOREBOARD            = 0x3D
	PACKET_CLIENT_TEAMS                         = 0x3E
	PACKET_CLIENT_PLUGIN_MESSAGE                = 0x3F
	PACKET_CLIENT_DISCONNECT                    = 0x40
	PACKET_CLIENT_DIFFICULTY                    = 0x41
	PACKET_CLIENT_COMBAT_EVENT                  = 0x42
	PACKET_CLIENT_CAMERA                        = 0x43
	PACKET_CLIENT_WORLD_BORDER                  = 0x44
	PACKET_CLIENT_TITLE                         = 0x45
	PACKET_CLIENT_SET_COMPRESSION               = 0x46
	PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT         = 0x47
	PACKET_CLIENT_RESOURCE_PACK                 = 0x48
	PACKET_CLIENT_UPDATE_ENTITY_NBT             = 0x49
	PACKET_SERVER_KEEPALIVE                     = 0x00
	PACKET_SERVER_CHAT                          = 0x01
	PACKET_SERVER_USE_ENTITY                    = 0x02
	PACKET_SERVER_PLAYER                        = 0x03
	PACKET_SERVER_PLAYER_POSITION               = 0x04
	PACKET_SERVER_PLAYER_LOOK                   = 0x05
	PACKET_SERVER_PLAYER_LOOK_AND_POSITION      = 0x06
	PACKET_SERVER_PLAYER_DIGGING                = 0x07
	PACKET_SERVER_PLAYER_BLOCK_PLACEMENT        = 0x08
	PACKET_SERVER_HELD_ITEM_CHANGE              = 0x09
	PACKET_SERVER_ANIMATION                     = 0x0A
	PACKET_SERVER_ENTITY_ACTION                 = 0x0B
	PACKET_SERVER_STEER_VEHICLE                 = 0x0C
	PACKET_SERVER_CLOSE_WINDOW                  = 0x0D
	PACKET_SERVER_CLICK_WINDOW                  = 0x0E
	PACKET_SERVER_CONFIRM_TRANSACTION           = 0x0F
	PACKET_SERVER_CREATIVE_INVENTORY_ACTION     = 0x10
	PACKET_SERVER_ENCHANT_ITEM                  = 0x11
	PACKET_SERVER_UPDATE_SIGN                   = 0x12
	PACKET_SERVER_PLAYER_ABILITIES              = 0x13
	PACKET_SERVER_TAB_COMPLETE                  = 0x14
	PACKET_SERVER_CLIENT_SETTINGS               = 0x15
	PACKET_SERVER_CLIENT_STATUS                 = 0x16
	PACKET_SERVER_PLUGIN_MESSAGE                = 0x17
	PACKET_SERVER_SPECTATE                      = 0x18
	PACKET_SERVER_RESOURCE_PACK_STATUS          = 0x19

	PACKET_CLIENT_LOGIN_DISCONNECT       = 0x00
	PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST  = 0x01
	PACKET_CLIENT_LOGIN_SUCCESS          = 0x02
	PACKET_CLIENT_LOGIN_SET_COMPRESSION  = 0x03
	PACKET_SERVER_LOGIN_START            = 0x00
	PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE = 0x01
)

var PlayPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	PACKET_CLIENT_KEEPALIVE:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_KEEPALIVE, Swappers),
	PACKET_CLIENT_JOIN_GAME:                     &CodecClientJoinGame{IdMap},
	PACKET_CLIENT_CHAT:                          minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHAT, Swappers),
	PACKET_CLIENT_TIME_UPDATE:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_TIME_UPDATE, Swappers),
	PACKET_CLIENT_ENTITY_EQUIPMENT:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_EQUIPMENT, Swappers),
	PACKET_CLIENT_SPAWN_POSITION:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_POSITION, Swappers),
	PACKET_CLIENT_UPDATE_HEALTH:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_HEALTH, Swappers),
	PACKET_CLIENT_RESPAWN:                       &mc18.CodecClientRespawn{IdMap},
	PACKET_CLIENT_PLAYER_POSITION_AND_LOOK:      minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLAYER_POSITION_AND_LOOK, Swappers),
	PACKET_CLIENT_HELD_ITEM_CHANGE:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_HELD_ITEM_CHANGE, Swappers),
	PACKET_CLIENT_USE_BED:                       minecraft.NewPacketGenericCodec(PACKET_CLIENT_USE_BED, Swappers),
	PACKET_CLIENT_ANIMATION:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_ANIMATION, Swappers),
	PACKET_CLIENT_SPAWN_PLAYER:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_PLAYER, Swappers),
	PACKET_CLIENT_COLLECT_ITEM:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_COLLECT_ITEM, Swappers),
	PACKET_CLIENT_SPAWN_OBJECT:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_OBJECT, Swappers),
	PACKET_CLIENT_SPAWN_MOB:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_MOB, Swappers),
	PACKET_CLIENT_SPAWN_PAINTING:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_PAINTING, Swappers),
	PACKET_CLIENT_SPAWN_EXPERIENCE_ORB:          minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_EXPERIENCE_ORB, Swappers),
	PACKET_CLIENT_ENTITY_VELOCITY:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_VELOCITY, Swappers),
	PACKET_CLIENT_DESTROY_ENTITIES:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_DESTROY_ENTITIES, Swappers),
	PACKET_CLIENT_ENTITY:                        minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY, Swappers),
	PACKET_CLIENT_ENTITY_RELATIVE_MOVE:          minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_RELATIVE_MOVE, Swappers),
	PACKET_CLIENT_ENTITY_LOOK:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_LOOK, Swappers),
	PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE: minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE, Swappers),
	PACKET_CLIENT_ENTITY_TELEPORT:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_TELEPORT, Swappers),
	PACKET_CLIENT_ENTITY_HEAD_LOOK:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_HEAD_LOOK, Swappers),
	PACKET_CLIENT_ENTITY_STATUS:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_STATUS, Swappers),
	PACKET_CLIENT_ATTACH_ENTITY:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_ATTACH_ENTITY, Swappers),
	PACKET_CLIENT_ENTITY_METADATA:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_METADATA, Swappers),
	PACKET_CLIENT_ENTITY_EFFECT:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_EFFECT, Swappers),
	PACKET_CLIENT_REMOVE_ENTITY_EFFECT:          minecraft.NewPacketGenericCodec(PACKET_CLIENT_REMOVE_ENTITY_EFFECT, Swappers),
	PACKET_CLIENT_SET_EXPERIENCE:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_EXPERIENCE, Swappers),
	PACKET_CLIENT_ENTITY_PROPERTIES:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_PROPERTIES, Swappers),
	PACKET_CLIENT_CHUNK_DATA:                    minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHUNK_DATA, Swappers),
	PACKET_CLIENT_MULTI_BLOCK_CHANGE:            minecraft.NewPacketGenericCodec(PACKET_CLIENT_MULTI_BLOCK_CHANGE, Swappers),
	PACKET_CLIENT_BLOCK_CHANGE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_BLOCK_CHANGE, Swappers),
	PACKET_CLIENT_BLOCK_ACTION:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_BLOCK_ACTION, Swappers),
	PACKET_CLIENT_BLOCK_BREAK_ANIMATION:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_BLOCK_BREAK_ANIMATION, Swappers),
	PACKET_CLIENT_MAP_CHUNK_BULK:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_MAP_CHUNK_BULK, Swappers),
	PACKET_CLIENT_EXPLOSION:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_EXPLOSION, Swappers),
	PACKET_CLIENT_EFFECT:                        minecraft.NewPacketGenericCodec(PACKET_CLIENT_EFFECT, Swappers),
	PACKET_CLIENT_NAMED_SOUND_EFFECT:            minecraft.NewPacketGenericCodec(PACKET_CLIENT_NAMED_SOUND_EFFECT, Swappers),
	PACKET_CLIENT_PARTICLE:                      minecraft.NewPacketGenericCodec(PACKET_CLIENT_PARTICLE, Swappers),
	PACKET_CLIENT_CHANGE_GAME_STATE:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHANGE_GAME_STATE, Swappers),
	PACKET_CLIENT_SPAWN_GLOBAL_ENTITY:           minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_GLOBAL_ENTITY, Swappers),
	PACKET_CLIENT_OPEN_WINDOW:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_OPEN_WINDOW, Swappers),
	PACKET_CLIENT_CLOSE_WINDOW:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_CLOSE_WINDOW, Swappers),
	PACKET_CLIENT_SET_SLOT:                      minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_SLOT, Swappers),
	PACKET_CLIENT_WINDOW_ITEMS:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_WINDOW_ITEMS, Swappers),
	PACKET_CLIENT_WINDOW_PROPERTY:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_WINDOW_PROPERTY, Swappers),
	PACKET_CLIENT_CONFIRM_TRANSACTION:           minecraft.NewPacketGenericCodec(PACKET_CLIENT_CONFIRM_TRANSACTION, Swappers),
	PACKET_CLIENT_UPDATE_SIGN:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_SIGN, Swappers),
	PACKET_CLIENT_MAPS:                          minecraft.NewPacketGenericCodec(PACKET_CLIENT_MAPS, Swappers),
	PACKET_CLIENT_UPDATE_BLOCK_ENTITY:           minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_BLOCK_ENTITY, Swappers),
	PACKET_CLIENT_SIGN_EDITOR_OPEN:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_SIGN_EDITOR_OPEN, Swappers),
	PACKET_CLIENT_STATISTICS:                    minecraft.NewPacketGenericCodec(PACKET_CLIENT_STATISTICS, Swappers),
	PACKET_CLIENT_PLAYER_LIST:                   &CodecClientPlayerList{},
	PACKET_CLIENT_PLAYER_ABILITIES:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLAYER_ABILITIES, Swappers),
	PACKET_CLIENT_TAB_COMPLETE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_TAB_COMPLETE, Swappers),
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE:          &CodecClientScoreboardObjective{IdMap},
	PACKET_CLIENT_UPDATE_SCORE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_SCORE, Swappers),
	PACKET_CLIENT_DISPLAY_SCOREBOARD:            minecraft.NewPacketGenericCodec(PACKET_CLIENT_DISPLAY_SCOREBOARD, Swappers),
	PACKET_CLIENT_TEAMS:                         &CodecClientTeams{IdMap},
	PACKET_CLIENT_PLUGIN_MESSAGE:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLUGIN_MESSAGE, Swappers),
	PACKET_CLIENT_DISCONNECT:                    &mc18.CodecClientDisconnect{IdMap},
}, []packet.PacketCodec{
	PACKET_SERVER_KEEPALIVE:                 minecraft.NewPacketGenericCodec(PACKET_SERVER_KEEPALIVE, Swappers),
	PACKET_SERVER_CHAT:                      minecraft.NewPacketGenericCodec(PACKET_SERVER_CHAT, Swappers),
	PACKET_SERVER_USE_ENTITY:                minecraft.NewPacketGenericCodec(PACKET_SERVER_USE_ENTITY, Swappers),
	PACKET_SERVER_PLAYER:                    minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER, Swappers),
	PACKET_SERVER_PLAYER_POSITION:           minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_POSITION, Swappers),
	PACKET_SERVER_PLAYER_LOOK:               minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_LOOK, Swappers),
	PACKET_SERVER_PLAYER_LOOK_AND_POSITION:  minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_LOOK_AND_POSITION, Swappers),
	PACKET_SERVER_PLAYER_DIGGING:            minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_DIGGING, Swappers),
	PACKET_SERVER_PLAYER_BLOCK_PLACEMENT:    minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_BLOCK_PLACEMENT, Swappers),
	PACKET_SERVER_HELD_ITEM_CHANGE:          minecraft.NewPacketGenericCodec(PACKET_SERVER_HELD_ITEM_CHANGE, Swappers),
	PACKET_SERVER_ANIMATION:                 minecraft.NewPacketGenericCodec(PACKET_SERVER_ANIMATION, Swappers),
	PACKET_SERVER_ENTITY_ACTION:             minecraft.NewPacketGenericCodec(PACKET_SERVER_ENTITY_ACTION, Swappers),
	PACKET_SERVER_STEER_VEHICLE:             minecraft.NewPacketGenericCodec(PACKET_SERVER_STEER_VEHICLE, Swappers),
	PACKET_SERVER_CLOSE_WINDOW:              minecraft.NewPacketGenericCodec(PACKET_SERVER_CLOSE_WINDOW, Swappers),
	PACKET_SERVER_CLICK_WINDOW:              minecraft.NewPacketGenericCodec(PACKET_SERVER_CLICK_WINDOW, Swappers),
	PACKET_SERVER_CONFIRM_TRANSACTION:       minecraft.NewPacketGenericCodec(PACKET_SERVER_CONFIRM_TRANSACTION, Swappers),
	PACKET_SERVER_CREATIVE_INVENTORY_ACTION: minecraft.NewPacketGenericCodec(PACKET_SERVER_CREATIVE_INVENTORY_ACTION, Swappers),
	PACKET_SERVER_ENCHANT_ITEM:              minecraft.NewPacketGenericCodec(PACKET_SERVER_ENCHANT_ITEM, Swappers),
	PACKET_SERVER_UPDATE_SIGN:               minecraft.NewPacketGenericCodec(PACKET_SERVER_UPDATE_SIGN, Swappers),
	PACKET_SERVER_PLAYER_ABILITIES:          minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_ABILITIES, Swappers),
	PACKET_SERVER_TAB_COMPLETE:              minecraft.NewPacketGenericCodec(PACKET_SERVER_TAB_COMPLETE, Swappers),
	PACKET_SERVER_CLIENT_SETTINGS:           &CodecServerClientSettings{IdMap},
	PACKET_SERVER_CLIENT_STATUS:             minecraft.NewPacketGenericCodec(PACKET_SERVER_CLIENT_STATUS, Swappers),
	PACKET_SERVER_PLUGIN_MESSAGE:            &CodecServerPluginMessage{IdMap},
})

var PlayPacketClientCodec = PlayPacketServerCodec.Flip()

var LoginPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	PACKET_CLIENT_LOGIN_DISCONNECT:      &mc18.CodecClientLoginDisconnect{IdMap},
	PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST: &CodecClientLoginEncryptRequest{IdMap},
	PACKET_CLIENT_LOGIN_SUCCESS:         &mc18.CodecClientLoginSuccess{IdMap},
}, []packet.PacketCodec{
	PACKET_SERVER_LOGIN_START:            &mc18.CodecServerLoginStart{IdMap},
	PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE: &CodecServerLoginEncryptResponse{IdMap},
})

var LoginPacketClientCodec = LoginPacketServerCodec.Flip()

var Swappers = &minecraft.PacketGenericSwappers{
	ClientInt: [][]int{
		PACKET_CLIENT_ENTITY_EQUIPMENT:              {0},
		PACKET_CLIENT_USE_BED:                       {0},
		PACKET_CLIENT_COLLECT_ITEM:                  {0, 4},
		PACKET_CLIENT_ENTITY_VELOCITY:               {0},
		PACKET_CLIENT_ENTITY:                        {0},
		PACKET_CLIENT_ENTITY_RELATIVE_MOVE:          {0},
		PACKET_CLIENT_ENTITY_LOOK:                   {0},
		PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE: {0},
		PACKET_CLIENT_ENTITY_TELEPORT:               {0},
		PACKET_CLIENT_ENTITY_HEAD_LOOK:              {0},
		PACKET_CLIENT_ENTITY_STATUS:                 {0},
		PACKET_CLIENT_ATTACH_ENTITY:                 {0, 4},
		PACKET_CLIENT_ENTITY_METADATA:               {0},
		PACKET_CLIENT_ENTITY_EFFECT:                 {0},
		PACKET_CLIENT_REMOVE_ENTITY_EFFECT:          {0},
		PACKET_CLIENT_ENTITY_PROPERTIES:             {0},
	},
	ClientVarInt: []bool{
		PACKET_CLIENT_ANIMATION:             true,
		PACKET_CLIENT_SPAWN_PLAYER:          true,
		PACKET_CLIENT_SPAWN_OBJECT:          true,
		PACKET_CLIENT_SPAWN_MOB:             true,
		PACKET_CLIENT_SPAWN_PAINTING:        true,
		PACKET_CLIENT_SPAWN_EXPERIENCE_ORB:  true,
		PACKET_CLIENT_BLOCK_BREAK_ANIMATION: true,
		PACKET_CLIENT_SPAWN_GLOBAL_ENTITY:   true,
	},
	ServerInt: [][]int{
		PACKET_SERVER_USE_ENTITY:    {0},
		PACKET_SERVER_ANIMATION:     {0},
		PACKET_SERVER_ENTITY_ACTION: {0},
	},
	ServerVarInt: []bool{},
}

var IdMap = mc18.IdMap

var Version = &minecraft.Version{
	Name:             "1.7",
	LoginClientCodec: LoginPacketClientCodec,
	LoginServerCodec: LoginPacketServerCodec,
	PlayClientCodec:  PlayPacketClientCodec,
	PlayServerCodec:  PlayPacketServerCodec,
	IdMap:            IdMap,
}

var VersionNum = 5
