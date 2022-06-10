package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
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

	ENTITY_ARROW          = 60
	ENTITY_FISHING_BOBBER = 90
	ENTITY_SPECTRAL_ARROW = -1
)

var PlayPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	PACKET_CLIENT_KEEPALIVE:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_KEEPALIVE, Swappers),
	PACKET_CLIENT_JOIN_GAME:                     &CodecClientJoinGame{IdMap},
	PACKET_CLIENT_CHAT:                          minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHAT, Swappers),
	PACKET_CLIENT_TIME_UPDATE:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_TIME_UPDATE, Swappers),
	PACKET_CLIENT_ENTITY_EQUIPMENT:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_EQUIPMENT, Swappers),
	PACKET_CLIENT_SPAWN_POSITION:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_POSITION, Swappers),
	PACKET_CLIENT_UPDATE_HEALTH:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_HEALTH, Swappers),
	PACKET_CLIENT_RESPAWN:                       &CodecClientRespawn{IdMap},
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
	PACKET_CLIENT_PLAYER_LIST:                   &CodecClientPlayerList{IdMap},
	PACKET_CLIENT_PLAYER_ABILITIES:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLAYER_ABILITIES, Swappers),
	PACKET_CLIENT_TAB_COMPLETE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_TAB_COMPLETE, Swappers),
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE:          &CodecClientScoreboardObjective{IdMap},
	PACKET_CLIENT_UPDATE_SCORE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_SCORE, Swappers),
	PACKET_CLIENT_DISPLAY_SCOREBOARD:            minecraft.NewPacketGenericCodec(PACKET_CLIENT_DISPLAY_SCOREBOARD, Swappers),
	PACKET_CLIENT_TEAMS:                         &CodecClientTeams{IdMap},
	PACKET_CLIENT_PLUGIN_MESSAGE:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLUGIN_MESSAGE, Swappers),
	PACKET_CLIENT_DISCONNECT:                    &CodecClientDisconnect{IdMap},
	PACKET_CLIENT_DIFFICULTY:                    minecraft.NewPacketGenericCodec(PACKET_CLIENT_DIFFICULTY, Swappers),
	PACKET_CLIENT_COMBAT_EVENT:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_COMBAT_EVENT, Swappers),
	PACKET_CLIENT_CAMERA:                        minecraft.NewPacketGenericCodec(PACKET_CLIENT_CAMERA, Swappers),
	PACKET_CLIENT_WORLD_BORDER:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_WORLD_BORDER, Swappers),
	PACKET_CLIENT_TITLE:                         minecraft.NewPacketGenericCodec(PACKET_CLIENT_TITLE, Swappers),
	PACKET_CLIENT_SET_COMPRESSION:               &CodecClientSetCompression{IdMap},
	PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT, Swappers),
	PACKET_CLIENT_RESOURCE_PACK:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_RESOURCE_PACK, Swappers),
	PACKET_CLIENT_UPDATE_ENTITY_NBT:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_ENTITY_NBT, Swappers),
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
	PACKET_SERVER_SPECTATE:                  minecraft.NewPacketGenericCodec(PACKET_SERVER_SPECTATE, Swappers),
	PACKET_SERVER_RESOURCE_PACK_STATUS:      minecraft.NewPacketGenericCodec(PACKET_SERVER_RESOURCE_PACK_STATUS, Swappers),
})

var PlayPacketClientCodec = PlayPacketServerCodec.Flip()

var LoginPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	PACKET_CLIENT_LOGIN_DISCONNECT:      &CodecClientLoginDisconnect{IdMap},
	PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST: &CodecClientLoginEncryptRequest{IdMap},
	PACKET_CLIENT_LOGIN_SUCCESS:         &CodecClientLoginSuccess{IdMap},
	PACKET_CLIENT_LOGIN_SET_COMPRESSION: &CodecClientLoginSetCompression{IdMap},
}, []packet.PacketCodec{
	PACKET_SERVER_LOGIN_START:            &CodecServerLoginStart{IdMap},
	PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE: &CodecServerLoginEncryptResponse{IdMap},
})

var LoginPacketClientCodec = LoginPacketServerCodec.Flip()

var Swappers = &minecraft.PacketGenericSwappers{
	ClientInt: [][]int{
		PACKET_CLIENT_ENTITY_STATUS: {0},
		PACKET_CLIENT_ATTACH_ENTITY: {0, 4},
	},
	ClientVarInt: []bool{
		PACKET_CLIENT_ENTITY_EQUIPMENT:              true,
		PACKET_CLIENT_USE_BED:                       true,
		PACKET_CLIENT_COLLECT_ITEM:                  true,
		PACKET_CLIENT_ANIMATION:                     true,
		PACKET_CLIENT_SPAWN_PLAYER:                  true,
		PACKET_CLIENT_SPAWN_OBJECT:                  true,
		PACKET_CLIENT_SPAWN_MOB:                     true,
		PACKET_CLIENT_SPAWN_PAINTING:                true,
		PACKET_CLIENT_SPAWN_EXPERIENCE_ORB:          true,
		PACKET_CLIENT_SPAWN_GLOBAL_ENTITY:           true,
		PACKET_CLIENT_ENTITY_VELOCITY:               true,
		PACKET_CLIENT_ENTITY:                        true,
		PACKET_CLIENT_ENTITY_RELATIVE_MOVE:          true,
		PACKET_CLIENT_ENTITY_LOOK:                   true,
		PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE: true,
		PACKET_CLIENT_ENTITY_TELEPORT:               true,
		PACKET_CLIENT_ENTITY_HEAD_LOOK:              true,
		PACKET_CLIENT_ENTITY_METADATA:               true,
		PACKET_CLIENT_ENTITY_EFFECT:                 true,
		PACKET_CLIENT_REMOVE_ENTITY_EFFECT:          true,
		PACKET_CLIENT_ENTITY_PROPERTIES:             true,
		PACKET_CLIENT_BLOCK_BREAK_ANIMATION:         true,
		PACKET_CLIENT_UPDATE_ENTITY_NBT:             true,
		PACKET_CLIENT_CAMERA:                        true,
	},
	ServerInt: [][]int{},
	ServerVarInt: []bool{
		PACKET_SERVER_USE_ENTITY:    true,
		PACKET_SERVER_ENTITY_ACTION: true,
	},
	IdMap: IdMap,
}

var IdMap = &minecraft.IdMap{
	PacketClientKeepalive:                 PACKET_CLIENT_KEEPALIVE,
	PacketClientJoinGame:                  PACKET_CLIENT_JOIN_GAME,
	PacketClientChat:                      PACKET_CLIENT_CHAT,
	PacketClientTimeUpdate:                PACKET_CLIENT_TIME_UPDATE,
	PacketClientEntityEquipment:           PACKET_CLIENT_ENTITY_EQUIPMENT,
	PacketClientSpawnPosition:             PACKET_CLIENT_SPAWN_POSITION,
	PacketClientUpdateHealth:              PACKET_CLIENT_UPDATE_HEALTH,
	PacketClientRespawn:                   PACKET_CLIENT_RESPAWN,
	PacketClientPlayerPositionandLook:     PACKET_CLIENT_PLAYER_POSITION_AND_LOOK,
	PacketClientHeldItemChange:            PACKET_CLIENT_HELD_ITEM_CHANGE,
	PacketClientUseBed:                    PACKET_CLIENT_USE_BED,
	PacketClientAnimation:                 PACKET_CLIENT_ANIMATION,
	PacketClientSpawnPlayer:               PACKET_CLIENT_SPAWN_PLAYER,
	PacketClientCollectItem:               PACKET_CLIENT_COLLECT_ITEM,
	PacketClientSpawnObject:               PACKET_CLIENT_SPAWN_OBJECT,
	PacketClientSpawnMob:                  PACKET_CLIENT_SPAWN_MOB,
	PacketClientSpawnPainting:             PACKET_CLIENT_SPAWN_PAINTING,
	PacketClientSpawnExperienceOrb:        PACKET_CLIENT_SPAWN_EXPERIENCE_ORB,
	PacketClientEntityVelocity:            PACKET_CLIENT_ENTITY_VELOCITY,
	PacketClientDestroyEntities:           PACKET_CLIENT_DESTROY_ENTITIES,
	PacketClientEntity:                    PACKET_CLIENT_ENTITY,
	PacketClientEntityRelativeMove:        PACKET_CLIENT_ENTITY_RELATIVE_MOVE,
	PacketClientEntityLook:                PACKET_CLIENT_ENTITY_LOOK,
	PacketClientEntityLookandRelativeMove: PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE,
	PacketClientEntityTeleport:            PACKET_CLIENT_ENTITY_TELEPORT,
	PacketClientEntityHeadLook:            PACKET_CLIENT_ENTITY_HEAD_LOOK,
	PacketClientEntityStatus:              PACKET_CLIENT_ENTITY_STATUS,
	PacketClientAttachEntity:              PACKET_CLIENT_ATTACH_ENTITY,
	PacketClientEntityMetadata:            PACKET_CLIENT_ENTITY_METADATA,
	PacketClientEntityEffect:              PACKET_CLIENT_ENTITY_EFFECT,
	PacketClientRemoveEntityEffect:        PACKET_CLIENT_REMOVE_ENTITY_EFFECT,
	PacketClientSetExperience:             PACKET_CLIENT_SET_EXPERIENCE,
	PacketClientEntityProperties:          PACKET_CLIENT_ENTITY_PROPERTIES,
	PacketClientChunkData:                 PACKET_CLIENT_CHUNK_DATA,
	PacketClientMultiBlockChange:          PACKET_CLIENT_MULTI_BLOCK_CHANGE,
	PacketClientBlockChange:               PACKET_CLIENT_BLOCK_CHANGE,
	PacketClientBlockAction:               PACKET_CLIENT_BLOCK_ACTION,
	PacketClientBlockBreakAnimation:       PACKET_CLIENT_BLOCK_BREAK_ANIMATION,
	PacketClientMapChunkBulk:              PACKET_CLIENT_MAP_CHUNK_BULK,
	PacketClientExplosion:                 PACKET_CLIENT_EXPLOSION,
	PacketClientEffect:                    PACKET_CLIENT_EFFECT,
	PacketClientNamedSoundEffect:          PACKET_CLIENT_NAMED_SOUND_EFFECT,
	PacketClientParticle:                  PACKET_CLIENT_PARTICLE,
	PacketClientChangeGameState:           PACKET_CLIENT_CHANGE_GAME_STATE,
	PacketClientSpawnGlobalEntity:         PACKET_CLIENT_SPAWN_GLOBAL_ENTITY,
	PacketClientOpenWindow:                PACKET_CLIENT_OPEN_WINDOW,
	PacketClientCloseWindow:               PACKET_CLIENT_CLOSE_WINDOW,
	PacketClientSetSlot:                   PACKET_CLIENT_SET_SLOT,
	PacketClientWindowItems:               PACKET_CLIENT_WINDOW_ITEMS,
	PacketClientWindowProperty:            PACKET_CLIENT_WINDOW_PROPERTY,
	PacketClientConfirmTransaction:        PACKET_CLIENT_CONFIRM_TRANSACTION,
	PacketClientUpdateSign:                PACKET_CLIENT_UPDATE_SIGN,
	PacketClientMaps:                      PACKET_CLIENT_MAPS,
	PacketClientUpdateBlockEntity:         PACKET_CLIENT_UPDATE_BLOCK_ENTITY,
	PacketClientSignEditorOpen:            PACKET_CLIENT_SIGN_EDITOR_OPEN,
	PacketClientStatistics:                PACKET_CLIENT_STATISTICS,
	PacketClientPlayerList:                PACKET_CLIENT_PLAYER_LIST,
	PacketClientPlayerAbilities:           PACKET_CLIENT_PLAYER_ABILITIES,
	PacketClientTabComplete:               PACKET_CLIENT_TAB_COMPLETE,
	PacketClientScoreboardObjective:       PACKET_CLIENT_SCOREBOARD_OBJECTIVE,
	PacketClientUpdateScore:               PACKET_CLIENT_UPDATE_SCORE,
	PacketClientDisplayScoreboard:         PACKET_CLIENT_DISPLAY_SCOREBOARD,
	PacketClientTeams:                     PACKET_CLIENT_TEAMS,
	PacketClientPluginMessage:             PACKET_CLIENT_PLUGIN_MESSAGE,
	PacketClientDisconnect:                PACKET_CLIENT_DISCONNECT,
	PacketClientDifficulty:                PACKET_CLIENT_DIFFICULTY,
	PacketClientCombatEvent:               PACKET_CLIENT_COMBAT_EVENT,
	PacketClientCamera:                    PACKET_CLIENT_CAMERA,
	PacketClientWorldBorder:               PACKET_CLIENT_WORLD_BORDER,
	PacketClientTitle:                     PACKET_CLIENT_TITLE,
	PacketClientSetCompression:            PACKET_CLIENT_SET_COMPRESSION,
	PacketClientPlayerListHeadFoot:        PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT,
	PacketClientResourcePack:              PACKET_CLIENT_RESOURCE_PACK,
	PacketClientUpdateEntityNbt:           PACKET_CLIENT_UPDATE_ENTITY_NBT,
	PacketServerKeepalive:                 PACKET_SERVER_KEEPALIVE,
	PacketServerChat:                      PACKET_SERVER_CHAT,
	PacketServerUseEntity:                 PACKET_SERVER_USE_ENTITY,
	PacketServerPlayer:                    PACKET_SERVER_PLAYER,
	PacketServerPlayerPosition:            PACKET_SERVER_PLAYER_POSITION,
	PacketServerPlayerLook:                PACKET_SERVER_PLAYER_LOOK,
	PacketServerPlayerLookandPosition:     PACKET_SERVER_PLAYER_LOOK_AND_POSITION,
	PacketServerPlayerDigging:             PACKET_SERVER_PLAYER_DIGGING,
	PacketServerPlayerBlockPlacement:      PACKET_SERVER_PLAYER_BLOCK_PLACEMENT,
	PacketServerHeldItemChange:            PACKET_SERVER_HELD_ITEM_CHANGE,
	PacketServerAnimation:                 PACKET_SERVER_ANIMATION,
	PacketServerEntityAction:              PACKET_SERVER_ENTITY_ACTION,
	PacketServerSteerVehicle:              PACKET_SERVER_STEER_VEHICLE,
	PacketServerCloseWindow:               PACKET_SERVER_CLOSE_WINDOW,
	PacketServerClickWindow:               PACKET_SERVER_CLICK_WINDOW,
	PacketServerConfirmTransaction:        PACKET_SERVER_CONFIRM_TRANSACTION,
	PacketServerCreativeInventoryAction:   PACKET_SERVER_CREATIVE_INVENTORY_ACTION,
	PacketServerEnchantItem:               PACKET_SERVER_ENCHANT_ITEM,
	PacketServerUpdateSign:                PACKET_SERVER_UPDATE_SIGN,
	PacketServerPlayerAbilities:           PACKET_SERVER_PLAYER_ABILITIES,
	PacketServerTabComplete:               PACKET_SERVER_TAB_COMPLETE,
	PacketServerClientSettings:            PACKET_SERVER_CLIENT_SETTINGS,
	PacketServerClientStatus:              PACKET_SERVER_CLIENT_STATUS,
	PacketServerPluginMessage:             PACKET_SERVER_PLUGIN_MESSAGE,
	PacketServerSpectate:                  PACKET_SERVER_SPECTATE,
	PacketServerResourcePackStatus:        PACKET_SERVER_RESOURCE_PACK_STATUS,
	// 1.9 - unsupported
	PacketClientBossBar:         -1,
	PacketClientSetCooldown:     -1,
	PacketClientUnloadChunk:     -1,
	PacketClientVehicleMove:     -1,
	PacketClientSetPassengers:   -1,
	PacketServerTeleportConfirm: -1,
	PacketServerVehicleMove:     -1,
	PacketServerSteerBoat:       -1,
	PacketServerUseItem:         -1,
	// 1.12 - unsupported
	PlayClientUnlockRecipes:       -1,
	PlayClientAdvancementProgress: -1,
	PlayClientAdvancements:        -1,
	PlayServerPrepareCraftingGrid: -1,
	PlayServerCraftingBookData:    -1,
	PlayServerAdvancementTab:      -1,
	// 1.13 - unsupported
	PacketClientStopSound: -1,
	// 1.14 - unsupported
	PacketClientUpdateViewDistance: -1,
	PacketClientEntitySoundEffect:  -1,
	// 1.18 - unsupported
	PacketClientUpdateSimulationDistance: -1,
	// 1.19 - unsupported
	PacketClientChatPreview:           -1,
	PacketClientPlayerChatMessage:     -1,
	PacketClientSystemChatMessage:     -1,
	PacketClientSetDisplayChatPreview: -1,
	PacketClientServerData:            -1,
	PacketServerChatCommand:           -1,
	PacketServerChatMessage:           -1,
	PacketServerChatPreview:           -1,
	PacketServerSetDisplayedRecipe:    -1,
	PacketServerSetBeaconEffect:       -1,

	PacketClientLoginDisconnect:      PACKET_CLIENT_LOGIN_DISCONNECT,
	PacketClientLoginEncryptRequest:  PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST,
	PacketClientLoginSuccess:         PACKET_CLIENT_LOGIN_SUCCESS,
	PacketClientLoginSetCompression:  PACKET_CLIENT_LOGIN_SET_COMPRESSION,
	PacketServerLoginStart:           PACKET_SERVER_LOGIN_START,
	PacketServerLoginEncryptResponse: PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE,

	EntityArrow:         ENTITY_ARROW,
	EntityFishingBobber: ENTITY_FISHING_BOBBER,
	EntitySpectralArrow: ENTITY_SPECTRAL_ARROW,
}

var Version = &minecraft.Version{
	Name:             "1.8",
	LoginClientCodec: LoginPacketClientCodec,
	LoginServerCodec: LoginPacketServerCodec,
	PlayClientCodec:  PlayPacketClientCodec,
	PlayServerCodec:  PlayPacketServerCodec,
	IdMap:            IdMap,
	Id: []int{
		47, // 1.8
	},
}
