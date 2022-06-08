package v119

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	mc113 "github.com/LilyPad/GoLilyPad/packet/minecraft/v113"
	mc114 "github.com/LilyPad/GoLilyPad/packet/minecraft/v114"
	mc116 "github.com/LilyPad/GoLilyPad/packet/minecraft/v116"
	mc1162 "github.com/LilyPad/GoLilyPad/packet/minecraft/v1162"
	mc118 "github.com/LilyPad/GoLilyPad/packet/minecraft/v118"
	mc18 "github.com/LilyPad/GoLilyPad/packet/minecraft/v18"
	mc19 "github.com/LilyPad/GoLilyPad/packet/minecraft/v19"
)

const (
	PACKET_CLIENT_SPAWN_OBJECT                  = 0x00
	PACKET_CLIENT_SPAWN_EXPERIENCE_ORB          = 0x01
	PACKET_CLIENT_SPAWN_PLAYER                  = 0x02
	PACKET_CLIENT_ANIMATION                     = 0x03
	PACKET_CLIENT_STATISTICS                    = 0x04
	PACKET_CLIENT_ACKNOWLEDGE_PLAYER_DIGGING    = 0x05
	PACKET_CLIENT_BLOCK_BREAK_ANIMATION         = 0x06
	PACKET_CLIENT_UPDATE_BLOCK_ENTITY           = 0x07
	PACKET_CLIENT_BLOCK_ACTION                  = 0x08
	PACKET_CLIENT_BLOCK_CHANGE                  = 0x09
	PACKET_CLIENT_BOSS_BAR                      = 0x0A
	PACKET_CLIENT_DIFFICULTY                    = 0x0B
	PACKET_CLIENT_CHAT_PREVIEW                  = 0x0C
	PACKET_CLIENT_CLEAR_TITLES                  = 0x0D
	PACKET_CLIENT_TAB_COMPLETE                  = 0x0E
	PACKET_CLIENT_DECLARE_COMMANDS              = 0x0F
	PACKET_CLIENT_CLOSE_WINDOW                  = 0x10
	PACKET_CLIENT_WINDOW_ITEMS                  = 0x11
	PACKET_CLIENT_WINDOW_PROPERTY               = 0x12
	PACKET_CLIENT_SET_SLOT                      = 0x13
	PACKET_CLIENT_SET_COOLDOWN                  = 0x14
	PACKET_CLIENT_PLUGIN_MESSAGE                = 0x15
	PACKET_CLIENT_NAMED_SOUND_EFFECT            = 0x16
	PACKET_CLIENT_DISCONNECT                    = 0x17
	PACKET_CLIENT_ENTITY_STATUS                 = 0x18
	PACKET_CLIENT_EXPLOSION                     = 0x19
	PACKET_CLIENT_UNLOAD_CHUNK                  = 0x1A
	PACKET_CLIENT_CHANGE_GAME_STATE             = 0x1B
	PACKET_CLIENT_OPEN_HORSE_WINDOW             = 0x1C
	PACKET_CLIENT_INITIALISE_WORLD_BORDER       = 0x1D
	PACKET_CLIENT_KEEPALIVE                     = 0x1E
	PACKET_CLIENT_CHUNK_DATA                    = 0x1F
	PACKET_CLIENT_EFFECT                        = 0x20
	PACKET_CLIENT_PARTICLE                      = 0x21
	PACKET_CLIENT_UPDATE_LIGHT                  = 0x22
	PACKET_CLIENT_JOIN_GAME                     = 0x23
	PACKET_CLIENT_MAPS                          = 0x24
	PACKET_CLIENT_TRADE_LIST                    = 0x25
	PACKET_CLIENT_ENTITY_RELATIVE_MOVE          = 0x26
	PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE = 0x27
	PACKET_CLIENT_ENTITY_LOOK                   = 0x28
	PACKET_CLIENT_VEHICLE_MOVE                  = 0x29
	PACKET_CLIENT_OPEN_BOOK                     = 0x2A
	PACKET_CLIENT_OPEN_WINDOW                   = 0x2B
	PACKET_CLIENT_SIGN_EDITOR_OPEN              = 0x2C
	PACKET_CLIENT_PING                          = 0x2D
	PACKET_CLIENT_CRAFT_RECIPE_RESPONSE         = 0x2E
	PACKET_CLIENT_PLAYER_ABILITIES              = 0x2F
	PACKET_CLIENT_PLAYER_CHAT_MESSAGE           = 0x30 // new
	PACKET_CLIENT_END_COMBAT_EVENT              = 0x31
	PACKET_CLIENT_ENTER_COMBAT_EVENT            = 0x32
	PACKET_CLIENT_DEATH_COMBAT_EVENT            = 0x33
	PACKET_CLIENT_PLAYER_LIST                   = 0x34
	PACKET_CLIENT_FACE_PLAYER                   = 0x35
	PACKET_CLIENT_PLAYER_POSITION_AND_LOOK      = 0x36
	PACKET_CLIENT_UNLOCK_RECIPES                = 0x37
	PACKET_CLIENT_DESTROY_ENTITIES              = 0x38
	PACKET_CLIENT_REMOVE_ENTITY_EFFECT          = 0x39
	PACKET_CLIENT_RESOURCE_PACK                 = 0x3A
	PACKET_CLIENT_RESPAWN                       = 0x3B
	PACKET_CLIENT_ENTITY_HEAD_LOOK              = 0x3C
	PACKET_CLIENT_MULTI_BLOCK_CHANGE            = 0x3D
	PACKET_CLIENT_ADVANCEMENT_PROGRESS          = 0x3E
	PACKET_CLIENT_SERVER_DATA                   = 0x3F
	PACKET_CLIENT_ACTION_BAR                    = 0x40
	PACKET_CLIENT_WORLD_BORDER_CENTER           = 0x41
	PACKET_CLIENT_WORLD_BORDER_LERP_SIZE        = 0x42
	PACKET_CLIENT_WORLD_BORDER_SIZE             = 0x43
	PACKET_CLIENT_WORLD_BORDER_WARNING_DELAY    = 0x44
	PACKET_CLIENT_WORLD_BORDER_WARNING_REACH    = 0x45
	PACKET_CLIENT_CAMERA                        = 0x46
	PACKET_CLIENT_HELD_ITEM_CHANGE              = 0x47
	PACKET_CLIENT_UPDATE_VIEW_POSITION          = 0x48
	PACKET_CLIENT_UPDATE_VIEW_DISTANCE          = 0x49
	PACKET_CLIENT_SPAWN_POSITION                = 0x4A
	PACKET_CLIENT_SET_DISPLAY_CHAT_PREVIEW      = 0x4B
	PACKET_CLIENT_DISPLAY_SCOREBOARD            = 0x4C
	PACKET_CLIENT_ENTITY_METADATA               = 0x4D
	PACKET_CLIENT_ATTACH_ENTITY                 = 0x4E
	PACKET_CLIENT_ENTITY_VELOCITY               = 0x4F
	PACKET_CLIENT_ENTITY_EQUIPMENT              = 0x50
	PACKET_CLIENT_SET_EXPERIENCE                = 0x51
	PACKET_CLIENT_UPDATE_HEALTH                 = 0x52
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE          = 0x53
	PACKET_CLIENT_SET_PASSENGERS                = 0x54
	PACKET_CLIENT_TEAMS                         = 0x55
	PACKET_CLIENT_UPDATE_SCORE                  = 0x56
	PACKET_CLIENT_UPDATE_SIMULATION_DISTANCE    = 0x57
	PACKET_CLIENT_SET_TITLE_SUBTITLE            = 0x58
	PACKET_CLIENT_TIME_UPDATE                   = 0x59
	PACKET_CLIENT_SET_TITLE_TEXT                = 0x5A
	PACKET_CLIENT_SET_TITLE_TIME                = 0x5B
	PACKET_CLIENT_ENTITY_SOUND_EFFECT           = 0x5C
	PACKET_CLIENT_SOUND_EFFECT                  = 0x5D
	PACKET_CLIENT_STOP_SOUND                    = 0x5E
	PACKET_CLIENT_SYSTEM_CHAT_MESSAGE           = 0x5F
	PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT         = 0x60
	PACKET_CLIENT_NBT_QUERY_RESPONSE            = 0x61
	PACKET_CLIENT_COLLECT_ITEM                  = 0x62
	PACKET_CLIENT_ENTITY_TELEPORT               = 0x63
	PACKET_CLIENT_ADVANCEMENTS                  = 0x64
	PACKET_CLIENT_ENTITY_PROPERTIES             = 0x65
	PACKET_CLIENT_ENTITY_EFFECT                 = 0x66
	PACKET_CLIENT_DECLARE_RECIPES               = 0x67
	PACKET_CLIENT_TAGS                          = 0x68
	PACKET_CLIENT_UPDATE_SIGN                   = -1
	PACKET_CLIENT_MAP_CHUNK_BULK                = -1
	PACKET_CLIENT_SET_COMPRESSION               = -1
	PACKET_CLIENT_UPDATE_ENTITY_NBT             = -1
	PACKET_CLIENT_USE_BED                       = -1
	PACKET_CLIENT_SPAWN_GLOBAL_ENTITY           = -1
	PACKET_CLIENT_CONFIRM_TRANSACTION           = -1
	PACKET_CLIENT_ENTITY                        = -1
	PACKET_CLIENT_COMBAT_EVENT                  = -1
	PACKET_CLIENT_WORLD_BORDER                  = -1
	PACKET_CLIENT_TITLE                         = -1
	PACKET_CLIENT_SPAWN_MOB                     = -1
	PACKET_CLIENT_SPAWN_PAINTING                = -1

	PACKET_SERVER_TELEPORT_CONFIRM              = 0x00
	PACKET_SERVER_QUERY_BLOCK_NBT               = 0x01
	PACKET_SERVER_SET_DIFFICULTY                = 0x02
	PACKET_SERVER_CHAT_COMMAND                  = 0x03
	PACKET_SERVER_CHAT_MESSAGE                  = 0x04
	PACKET_SERVER_CHAT_PREVIEW                  = 0x05
	PACKET_SERVER_CLIENT_STATUS                 = 0x06
	PACKET_SERVER_CLIENT_SETTINGS               = 0x07
	PACKET_SERVER_TAB_COMPLETE                  = 0x08
	PACKET_SERVER_CONFIRM_TRANSACTION           = 0x09
	PACKET_SERVER_CLICK_WINDOW                  = 0x0A
	PACKET_SERVER_CLOSE_WINDOW                  = 0x0B
	PACKET_SERVER_PLUGIN_MESSAGE                = 0x0C
	PACKET_SERVER_EDIT_BOOK                     = 0x0D
	PACKET_SERVER_QUERY_ENTITY_NBT              = 0x0E
	PACKET_SERVER_USE_ENTITY                    = 0x0F
	PACKET_SERVER_GENERATE_STRUCTURE            = 0x10
	PACKET_SERVER_KEEPALIVE                     = 0x11
	PACKET_SERVER_LOCK_DIFFICULTY               = 0x12
	PACKET_SERVER_PLAYER_POSITION               = 0x13
	PACKET_SERVER_PLAYER_LOOK_AND_POSITION      = 0x14
	PACKET_SERVER_PLAYER_LOOK                   = 0x15
	PACKET_SERVER_PLAYER                        = 0x16
	PACKET_SERVER_VEHICLE_MOVE                  = 0x17
	PACKET_SERVER_STEER_BOAT                    = 0x18
	PACKET_SERVER_PICK_ITEM                     = 0x19
	PACKET_SERVER_CRAFT_RECIPE_REQUEST          = 0x1A
	PACKET_SERVER_PLAYER_ABILITIES              = 0x1B
	PACKET_SERVER_PLAYER_DIGGING                = 0x1C
	PACKET_SERVER_ENTITY_ACTION                 = 0x1D
	PACKET_SERVER_STEER_VEHICLE                 = 0x1E
	PACKET_SERVER_PONG                          = 0x1F
	PACKET_SERVER_SET_RECIPE_BOOK_STATE         = 0x20
	PACKET_SERVER_SET_DISPLAYED_RECIPE          = 0x21
	PACKET_SERVER_NAME_ITEM                     = 0x22
	PACKET_SERVER_RESOURCE_PACK_STATUS          = 0x23
	PACKET_SERVER_ADVANCEMENT_TAB               = 0x24
	PACKET_SERVER_SELECT_TRADE                  = 0x25
	PACKET_SERVER_SET_BEACON_EFFECT             = 0x26
	PACKET_SERVER_HELD_ITEM_CHANGE              = 0x27
	PACKET_SERVER_UPDATE_COMMAND_BLOCK          = 0x28
	PACKET_SERVER_UPDATE_COMMAND_BLOCK_MINECART = 0x29
	PACKET_SERVER_CREATIVE_INVENTORY_ACTION     = 0x2A
	PACKET_SERVER_UPDATE_JIGSAW_BLOCK           = 0x2B
	PACKET_SERVER_UPDATE_STRUCTURE_BLOCK        = 0x2C
	PACKET_SERVER_UPDATE_SIGN                   = 0x2D
	PACKET_SERVER_ANIMATION                     = 0x2E
	PACKET_SERVER_SPECTATE                      = 0x2F
	PACKET_SERVER_PLAYER_BLOCK_PLACEMENT        = 0x30
	PACKET_SERVER_USE_ITEM                      = 0x31
	PACKET_SERVER_PREPARE_CRAFTING_GRID         = -1
	PACKET_SERVER_ENCHANT_ITEM                  = -1

	ENTITY_ARROW          = 2
	ENTITY_FISHING_BOBBER = 102
	ENTITY_SPECTRAL_ARROW = 72
)

var PlayPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	PACKET_CLIENT_KEEPALIVE:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_KEEPALIVE, Swappers),
	PACKET_CLIENT_JOIN_GAME:                     &mc118.CodecClientJoinGame{IdMap},
	PACKET_CLIENT_TIME_UPDATE:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_TIME_UPDATE, Swappers),
	PACKET_CLIENT_ENTITY_EQUIPMENT:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_EQUIPMENT, Swappers),
	PACKET_CLIENT_SPAWN_POSITION:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_POSITION, Swappers),
	PACKET_CLIENT_UPDATE_HEALTH:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_HEALTH, Swappers),
	PACKET_CLIENT_RESPAWN:                       &mc1162.CodecClientRespawn{IdMap},
	PACKET_CLIENT_PLAYER_POSITION_AND_LOOK:      minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLAYER_POSITION_AND_LOOK, Swappers),
	PACKET_CLIENT_HELD_ITEM_CHANGE:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_HELD_ITEM_CHANGE, Swappers),
	PACKET_CLIENT_ANIMATION:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_ANIMATION, Swappers),
	PACKET_CLIENT_SPAWN_PLAYER:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_PLAYER, Swappers),
	PACKET_CLIENT_COLLECT_ITEM:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_COLLECT_ITEM, Swappers),
	PACKET_CLIENT_SPAWN_OBJECT:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_OBJECT, Swappers),
	PACKET_CLIENT_SPAWN_EXPERIENCE_ORB:          minecraft.NewPacketGenericCodec(PACKET_CLIENT_SPAWN_EXPERIENCE_ORB, Swappers),
	PACKET_CLIENT_ENTITY_VELOCITY:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_VELOCITY, Swappers),
	PACKET_CLIENT_DESTROY_ENTITIES:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_DESTROY_ENTITIES, Swappers),
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
	PACKET_CLIENT_EXPLOSION:                     minecraft.NewPacketGenericCodec(PACKET_CLIENT_EXPLOSION, Swappers),
	PACKET_CLIENT_EFFECT:                        minecraft.NewPacketGenericCodec(PACKET_CLIENT_EFFECT, Swappers),
	PACKET_CLIENT_NAMED_SOUND_EFFECT:            minecraft.NewPacketGenericCodec(PACKET_CLIENT_NAMED_SOUND_EFFECT, Swappers),
	PACKET_CLIENT_PARTICLE:                      minecraft.NewPacketGenericCodec(PACKET_CLIENT_PARTICLE, Swappers),
	PACKET_CLIENT_CHANGE_GAME_STATE:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHANGE_GAME_STATE, Swappers),
	PACKET_CLIENT_OPEN_WINDOW:                   minecraft.NewPacketGenericCodec(PACKET_CLIENT_OPEN_WINDOW, Swappers),
	PACKET_CLIENT_CLOSE_WINDOW:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_CLOSE_WINDOW, Swappers),
	PACKET_CLIENT_SET_SLOT:                      minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_SLOT, Swappers),
	PACKET_CLIENT_WINDOW_ITEMS:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_WINDOW_ITEMS, Swappers),
	PACKET_CLIENT_WINDOW_PROPERTY:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_WINDOW_PROPERTY, Swappers),
	PACKET_CLIENT_MAPS:                          minecraft.NewPacketGenericCodec(PACKET_CLIENT_MAPS, Swappers),
	PACKET_CLIENT_UPDATE_BLOCK_ENTITY:           minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_BLOCK_ENTITY, Swappers),
	PACKET_CLIENT_SIGN_EDITOR_OPEN:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_SIGN_EDITOR_OPEN, Swappers),
	PACKET_CLIENT_PING:                          minecraft.NewPacketGenericCodec(PACKET_CLIENT_PING, Swappers),
	PACKET_CLIENT_STATISTICS:                    minecraft.NewPacketGenericCodec(PACKET_CLIENT_STATISTICS, Swappers),
	PACKET_CLIENT_PLAYER_LIST:                   &mc18.CodecClientPlayerList{IdMap},
	PACKET_CLIENT_TAB_COMPLETE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_TAB_COMPLETE, Swappers),
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE:          &mc113.CodecClientScoreboardObjective{IdMap},
	PACKET_CLIENT_UPDATE_SCORE:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_SCORE, Swappers),
	PACKET_CLIENT_DISPLAY_SCOREBOARD:            minecraft.NewPacketGenericCodec(PACKET_CLIENT_DISPLAY_SCOREBOARD, Swappers),
	PACKET_CLIENT_TEAMS:                         &mc113.CodecClientTeams{IdMap},
	PACKET_CLIENT_PLUGIN_MESSAGE:                minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLUGIN_MESSAGE, Swappers),
	PACKET_CLIENT_DISCONNECT:                    &mc18.CodecClientDisconnect{IdMap},
	PACKET_CLIENT_DIFFICULTY:                    minecraft.NewPacketGenericCodec(PACKET_CLIENT_DIFFICULTY, Swappers),
	PACKET_CLIENT_CAMERA:                        minecraft.NewPacketGenericCodec(PACKET_CLIENT_CAMERA, Swappers),
	PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_PLAYER_LIST_HEAD_FOOT, Swappers),
	PACKET_CLIENT_RESOURCE_PACK:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_RESOURCE_PACK, Swappers),
	// 1.9
	PACKET_CLIENT_BOSS_BAR:       &mc19.CodecClientBossBar{IdMap},
	PACKET_CLIENT_SET_COOLDOWN:   minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_COOLDOWN, Swappers),
	PACKET_CLIENT_UNLOAD_CHUNK:   minecraft.NewPacketGenericCodec(PACKET_CLIENT_UNLOAD_CHUNK, Swappers),
	PACKET_CLIENT_VEHICLE_MOVE:   minecraft.NewPacketGenericCodec(PACKET_CLIENT_VEHICLE_MOVE, Swappers),
	PACKET_CLIENT_SET_PASSENGERS: minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_PASSENGERS, Swappers),
	PACKET_CLIENT_SOUND_EFFECT:   minecraft.NewPacketGenericCodec(PACKET_CLIENT_SOUND_EFFECT, Swappers),
	// 1.12
	PACKET_CLIENT_UNLOCK_RECIPES:       minecraft.NewPacketGenericCodec(PACKET_CLIENT_UNLOCK_RECIPES, Swappers),
	PACKET_CLIENT_ADVANCEMENT_PROGRESS: minecraft.NewPacketGenericCodec(PACKET_CLIENT_ADVANCEMENT_PROGRESS, Swappers),
	PACKET_CLIENT_ADVANCEMENTS:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_ADVANCEMENTS, Swappers),
	// 1.12.1
	PACKET_CLIENT_CRAFT_RECIPE_RESPONSE: minecraft.NewPacketGenericCodec(PACKET_CLIENT_CRAFT_RECIPE_RESPONSE, Swappers),
	// 1.13
	PACKET_CLIENT_DECLARE_COMMANDS:   minecraft.NewPacketGenericCodec(PACKET_CLIENT_DECLARE_COMMANDS, Swappers),
	PACKET_CLIENT_NBT_QUERY_RESPONSE: minecraft.NewPacketGenericCodec(PACKET_CLIENT_NBT_QUERY_RESPONSE, Swappers),
	PACKET_CLIENT_FACE_PLAYER:        minecraft.NewPacketGenericCodec(PACKET_CLIENT_FACE_PLAYER, Swappers),
	PACKET_CLIENT_DECLARE_RECIPES:    minecraft.NewPacketGenericCodec(PACKET_CLIENT_DECLARE_RECIPES, Swappers),
	PACKET_CLIENT_TAGS:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_TAGS, Swappers),
	PACKET_CLIENT_STOP_SOUND:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_STOP_SOUND, Swappers),
	// 1.14
	PACKET_CLIENT_OPEN_HORSE_WINDOW:          minecraft.NewPacketGenericCodec(PACKET_CLIENT_OPEN_HORSE_WINDOW, Swappers),
	PACKET_CLIENT_UPDATE_LIGHT:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_LIGHT, Swappers),
	PACKET_CLIENT_TRADE_LIST:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_TRADE_LIST, Swappers),
	PACKET_CLIENT_OPEN_BOOK:                  minecraft.NewPacketGenericCodec(PACKET_CLIENT_OPEN_BOOK, Swappers),
	PACKET_CLIENT_UPDATE_VIEW_POSITION:       minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_VIEW_POSITION, Swappers),
	PACKET_CLIENT_UPDATE_VIEW_DISTANCE:       &mc114.CodecClientViewDistance{IdMap},
	PACKET_CLIENT_ENTITY_SOUND_EFFECT:        minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTITY_SOUND_EFFECT, Swappers),
	PACKET_CLIENT_ACKNOWLEDGE_PLAYER_DIGGING: minecraft.NewPacketGenericCodec(PACKET_CLIENT_ACKNOWLEDGE_PLAYER_DIGGING, Swappers),
	// 1.17
	PACKET_CLIENT_CLEAR_TITLES:               minecraft.NewPacketGenericCodec(PACKET_CLIENT_CLEAR_TITLES, Swappers),
	PACKET_CLIENT_INITIALISE_WORLD_BORDER:    minecraft.NewPacketGenericCodec(PACKET_CLIENT_INITIALISE_WORLD_BORDER, Swappers),
	PACKET_CLIENT_END_COMBAT_EVENT:           minecraft.NewPacketGenericCodec(PACKET_CLIENT_END_COMBAT_EVENT, Swappers),
	PACKET_CLIENT_ENTER_COMBAT_EVENT:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_ENTER_COMBAT_EVENT, Swappers),
	PACKET_CLIENT_DEATH_COMBAT_EVENT:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_DEATH_COMBAT_EVENT, Swappers),
	PACKET_CLIENT_ACTION_BAR:                 minecraft.NewPacketGenericCodec(PACKET_CLIENT_ACTION_BAR, Swappers),
	PACKET_CLIENT_WORLD_BORDER_CENTER:        minecraft.NewPacketGenericCodec(PACKET_CLIENT_WORLD_BORDER_CENTER, Swappers),
	PACKET_CLIENT_WORLD_BORDER_LERP_SIZE:     minecraft.NewPacketGenericCodec(PACKET_CLIENT_WORLD_BORDER_LERP_SIZE, Swappers),
	PACKET_CLIENT_WORLD_BORDER_SIZE:          minecraft.NewPacketGenericCodec(PACKET_CLIENT_WORLD_BORDER_SIZE, Swappers),
	PACKET_CLIENT_WORLD_BORDER_WARNING_DELAY: minecraft.NewPacketGenericCodec(PACKET_CLIENT_WORLD_BORDER_WARNING_DELAY, Swappers),
	PACKET_CLIENT_WORLD_BORDER_WARNING_REACH: minecraft.NewPacketGenericCodec(PACKET_CLIENT_WORLD_BORDER_WARNING_REACH, Swappers),
	PACKET_CLIENT_SET_TITLE_SUBTITLE:         minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_TITLE_SUBTITLE, Swappers),
	PACKET_CLIENT_SET_TITLE_TEXT:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_TITLE_TEXT, Swappers),
	PACKET_CLIENT_SET_TITLE_TIME:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_TITLE_TIME, Swappers),
	// 1.18
	PACKET_CLIENT_UPDATE_SIMULATION_DISTANCE: minecraft.NewPacketGenericCodec(PACKET_CLIENT_UPDATE_SIMULATION_DISTANCE, Swappers),
	// 1.19
	PACKET_CLIENT_CHAT_PREVIEW:             minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHAT_PREVIEW, Swappers),
	PACKET_CLIENT_PLAYER_CHAT_MESSAGE:      minecraft.NewPacketGenericCodec(PACKET_CLIENT_CHAT_PREVIEW, Swappers),
	PACKET_CLIENT_SERVER_DATA:              minecraft.NewPacketGenericCodec(PACKET_CLIENT_SERVER_DATA, Swappers),
	PACKET_CLIENT_SET_DISPLAY_CHAT_PREVIEW: minecraft.NewPacketGenericCodec(PACKET_CLIENT_SET_DISPLAY_CHAT_PREVIEW, Swappers),
	PACKET_CLIENT_SYSTEM_CHAT_MESSAGE:      minecraft.NewPacketGenericCodec(PACKET_CLIENT_SYSTEM_CHAT_MESSAGE, Swappers),
}, []packet.PacketCodec{
	PACKET_SERVER_KEEPALIVE:                 minecraft.NewPacketGenericCodec(PACKET_SERVER_KEEPALIVE, Swappers),
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
	PACKET_SERVER_UPDATE_SIGN:               minecraft.NewPacketGenericCodec(PACKET_SERVER_UPDATE_SIGN, Swappers),
	PACKET_SERVER_PLAYER_ABILITIES:          minecraft.NewPacketGenericCodec(PACKET_SERVER_PLAYER_ABILITIES, Swappers),
	PACKET_SERVER_TAB_COMPLETE:              minecraft.NewPacketGenericCodec(PACKET_SERVER_TAB_COMPLETE, Swappers),
	PACKET_SERVER_CLIENT_SETTINGS:           &mc118.CodecServerClientSettings{IdMap},
	PACKET_SERVER_CLIENT_STATUS:             minecraft.NewPacketGenericCodec(PACKET_SERVER_CLIENT_STATUS, Swappers),
	PACKET_SERVER_PLUGIN_MESSAGE:            &mc18.CodecServerPluginMessage{IdMap},
	PACKET_SERVER_SPECTATE:                  minecraft.NewPacketGenericCodec(PACKET_SERVER_SPECTATE, Swappers),
	PACKET_SERVER_RESOURCE_PACK_STATUS:      minecraft.NewPacketGenericCodec(PACKET_SERVER_RESOURCE_PACK_STATUS, Swappers),
	// 1.9
	PACKET_SERVER_TELEPORT_CONFIRM: minecraft.NewPacketGenericCodec(PACKET_SERVER_TELEPORT_CONFIRM, Swappers),
	PACKET_SERVER_VEHICLE_MOVE:     minecraft.NewPacketGenericCodec(PACKET_SERVER_VEHICLE_MOVE, Swappers),
	PACKET_SERVER_STEER_BOAT:       minecraft.NewPacketGenericCodec(PACKET_SERVER_STEER_BOAT, Swappers),
	PACKET_SERVER_USE_ITEM:         minecraft.NewPacketGenericCodec(PACKET_SERVER_USE_ITEM, Swappers),
	// 1.12
	PACKET_SERVER_ADVANCEMENT_TAB: minecraft.NewPacketGenericCodec(PACKET_SERVER_ADVANCEMENT_TAB, Swappers),
	// 1.12.1
	PACKET_SERVER_CRAFT_RECIPE_REQUEST: minecraft.NewPacketGenericCodec(PACKET_SERVER_CRAFT_RECIPE_REQUEST, Swappers),
	// 1.13
	PACKET_SERVER_QUERY_BLOCK_NBT:               minecraft.NewPacketGenericCodec(PACKET_SERVER_QUERY_BLOCK_NBT, Swappers),
	PACKET_SERVER_EDIT_BOOK:                     minecraft.NewPacketGenericCodec(PACKET_SERVER_EDIT_BOOK, Swappers),
	PACKET_SERVER_QUERY_ENTITY_NBT:              minecraft.NewPacketGenericCodec(PACKET_SERVER_QUERY_ENTITY_NBT, Swappers),
	PACKET_SERVER_PICK_ITEM:                     minecraft.NewPacketGenericCodec(PACKET_SERVER_PICK_ITEM, Swappers),
	PACKET_SERVER_NAME_ITEM:                     minecraft.NewPacketGenericCodec(PACKET_SERVER_NAME_ITEM, Swappers),
	PACKET_SERVER_SELECT_TRADE:                  minecraft.NewPacketGenericCodec(PACKET_SERVER_SELECT_TRADE, Swappers),
	PACKET_SERVER_UPDATE_COMMAND_BLOCK:          minecraft.NewPacketGenericCodec(PACKET_SERVER_UPDATE_COMMAND_BLOCK, Swappers),
	PACKET_SERVER_UPDATE_COMMAND_BLOCK_MINECART: minecraft.NewPacketGenericCodec(PACKET_SERVER_UPDATE_COMMAND_BLOCK_MINECART, Swappers),
	PACKET_SERVER_UPDATE_STRUCTURE_BLOCK:        minecraft.NewPacketGenericCodec(PACKET_SERVER_UPDATE_STRUCTURE_BLOCK, Swappers),
	// 1.14
	PACKET_SERVER_SET_DIFFICULTY:      minecraft.NewPacketGenericCodec(PACKET_SERVER_SET_DIFFICULTY, Swappers),
	PACKET_SERVER_LOCK_DIFFICULTY:     minecraft.NewPacketGenericCodec(PACKET_SERVER_LOCK_DIFFICULTY, Swappers),
	PACKET_SERVER_UPDATE_JIGSAW_BLOCK: minecraft.NewPacketGenericCodec(PACKET_SERVER_UPDATE_JIGSAW_BLOCK, Swappers),
	// 1.16
	PACKET_SERVER_GENERATE_STRUCTURE: minecraft.NewPacketGenericCodec(PACKET_SERVER_GENERATE_STRUCTURE, Swappers),
	// 1.17
	PACKET_SERVER_PONG: minecraft.NewPacketGenericCodec(PACKET_SERVER_PONG, Swappers),
	// 1.19
	PACKET_SERVER_CHAT_COMMAND:          minecraft.NewPacketGenericCodec(PACKET_SERVER_CHAT_COMMAND, Swappers),
	PACKET_SERVER_CHAT_MESSAGE:          minecraft.NewPacketGenericCodec(PACKET_SERVER_CHAT_MESSAGE, Swappers),
	PACKET_SERVER_CHAT_PREVIEW:          minecraft.NewPacketGenericCodec(PACKET_SERVER_CHAT_PREVIEW, Swappers),
	PACKET_SERVER_SET_RECIPE_BOOK_STATE: minecraft.NewPacketGenericCodec(PACKET_SERVER_SET_RECIPE_BOOK_STATE, Swappers),
	PACKET_SERVER_SET_DISPLAYED_RECIPE:  minecraft.NewPacketGenericCodec(PACKET_SERVER_SET_DISPLAYED_RECIPE, Swappers),
	PACKET_SERVER_SET_BEACON_EFFECT:     minecraft.NewPacketGenericCodec(PACKET_SERVER_SET_BEACON_EFFECT, Swappers),
})

var PlayPacketClientCodec = PlayPacketServerCodec.Flip()

var LoginPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	mc18.PACKET_CLIENT_LOGIN_DISCONNECT:      &mc18.CodecClientLoginDisconnect{IdMap},
	mc18.PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST: &mc18.CodecClientLoginEncryptRequest{IdMap},
	mc18.PACKET_CLIENT_LOGIN_SUCCESS:         &mc116.CodecClientLoginSuccess{IdMap},
	mc18.PACKET_CLIENT_LOGIN_SET_COMPRESSION: &mc18.CodecClientLoginSetCompression{IdMap},
}, []packet.PacketCodec{
	mc18.PACKET_SERVER_LOGIN_START:            &mc18.CodecServerLoginStart{IdMap},
	mc18.PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE: &mc18.CodecServerLoginEncryptResponse{IdMap},
})

var LoginPacketClientCodec = LoginPacketServerCodec.Flip()

var Swappers = &minecraft.PacketGenericSwappers{
	ClientInt: [][]int{
		PACKET_CLIENT_ENTITY_STATUS: {0},
		PACKET_CLIENT_ATTACH_ENTITY: {0, 4},
	},
	ClientVarInt: []bool{
		PACKET_CLIENT_SPAWN_OBJECT:                  true,
		PACKET_CLIENT_SPAWN_PLAYER:                  true,
		PACKET_CLIENT_ENTITY_EQUIPMENT:              true,
		PACKET_CLIENT_SET_PASSENGERS:                true,
		PACKET_CLIENT_ENTITY_PROPERTIES:             true,
		PACKET_CLIENT_COLLECT_ITEM:                  true,
		PACKET_CLIENT_ANIMATION:                     true,
		PACKET_CLIENT_SPAWN_EXPERIENCE_ORB:          true,
		PACKET_CLIENT_ENTITY_VELOCITY:               true,
		PACKET_CLIENT_ENTITY_RELATIVE_MOVE:          true,
		PACKET_CLIENT_ENTITY_LOOK:                   true,
		PACKET_CLIENT_ENTITY_LOOK_AND_RELATIVE_MOVE: true,
		PACKET_CLIENT_ENTITY_TELEPORT:               true,
		PACKET_CLIENT_ENTITY_HEAD_LOOK:              true,
		PACKET_CLIENT_ENTITY_METADATA:               true,
		PACKET_CLIENT_ENTITY_EFFECT:                 true,
		PACKET_CLIENT_REMOVE_ENTITY_EFFECT:          true,
		PACKET_CLIENT_BLOCK_BREAK_ANIMATION:         true,
		PACKET_CLIENT_CAMERA:                        true,
	},
	ClientSpawnObject: minecraft.PacketGenericSwappersClientSpawnObjectUUIDTypeVar,
	ServerInt:         [][]int{},
	ServerVarInt: []bool{
		PACKET_SERVER_ENTITY_ACTION: true,
		PACKET_SERVER_USE_ENTITY:    true,
	},
	IdMap: IdMap,
}

var IdMap = &minecraft.IdMap{
	PacketClientKeepalive:                 PACKET_CLIENT_KEEPALIVE,
	PacketClientJoinGame:                  PACKET_CLIENT_JOIN_GAME,
	PacketClientChat:                      -1,
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
	PacketServerChat:                      -1,
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
	// 1.9
	PacketClientBossBar:         PACKET_CLIENT_BOSS_BAR,
	PacketClientSetCooldown:     PACKET_CLIENT_SET_COOLDOWN,
	PacketClientUnloadChunk:     PACKET_CLIENT_UNLOAD_CHUNK,
	PacketClientVehicleMove:     PACKET_CLIENT_VEHICLE_MOVE,
	PacketClientSetPassengers:   PACKET_CLIENT_SET_PASSENGERS,
	PacketServerTeleportConfirm: PACKET_SERVER_TELEPORT_CONFIRM,
	PacketServerVehicleMove:     PACKET_SERVER_VEHICLE_MOVE,
	PacketServerSteerBoat:       PACKET_SERVER_STEER_BOAT,
	PacketServerUseItem:         PACKET_SERVER_USE_ITEM,
	// 1.12
	PlayClientUnlockRecipes:       PACKET_CLIENT_UNLOCK_RECIPES,
	PlayClientAdvancementProgress: PACKET_CLIENT_ADVANCEMENT_PROGRESS,
	PlayClientAdvancements:        PACKET_CLIENT_ADVANCEMENTS,
	PlayServerPrepareCraftingGrid: PACKET_SERVER_PREPARE_CRAFTING_GRID,
	PlayServerCraftingBookData:    PACKET_SERVER_SET_RECIPE_BOOK_STATE,
	PlayServerAdvancementTab:      PACKET_SERVER_ADVANCEMENT_TAB,
	// 1.13
	PacketClientStopSound: PACKET_CLIENT_STOP_SOUND,
	// 1.14
	PacketClientUpdateViewDistance: PACKET_CLIENT_UPDATE_VIEW_DISTANCE,
	PacketClientEntitySoundEffect:  PACKET_CLIENT_ENTITY_SOUND_EFFECT,
	// 1.18
	PacketClientUpdateSimulationDistance: PACKET_CLIENT_UPDATE_SIMULATION_DISTANCE,
	// 1.19
	PacketClientChatPreview:           PACKET_CLIENT_CHAT_PREVIEW,
	PacketClientPlayerChatMessage:     PACKET_CLIENT_PLAYER_CHAT_MESSAGE,
	PacketClientSystemChatMessage:     PACKET_CLIENT_SYSTEM_CHAT_MESSAGE,
	PacketClientSetDisplayChatPreview: PACKET_CLIENT_SET_DISPLAY_CHAT_PREVIEW,
	PacketClientServerData:            PACKET_CLIENT_SERVER_DATA,
	PacketServerChatCommand:           PACKET_SERVER_CHAT_COMMAND,
	PacketServerChatMessage:           PACKET_SERVER_CHAT_MESSAGE,
	PacketServerChatPreview:           PACKET_SERVER_CHAT_PREVIEW,
	PacketServerSetDisplayedRecipe:    PACKET_SERVER_SET_DISPLAYED_RECIPE,

	PacketClientLoginDisconnect:      mc18.PACKET_CLIENT_LOGIN_DISCONNECT,
	PacketClientLoginEncryptRequest:  mc18.PACKET_CLIENT_LOGIN_ENCRYPT_REQUEST,
	PacketClientLoginSuccess:         mc18.PACKET_CLIENT_LOGIN_SUCCESS,
	PacketClientLoginSetCompression:  mc18.PACKET_CLIENT_LOGIN_SET_COMPRESSION,
	PacketServerLoginStart:           mc18.PACKET_SERVER_LOGIN_START,
	PacketServerLoginEncryptResponse: mc18.PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE,

	EntityArrow:         ENTITY_ARROW,
	EntityFishingBobber: ENTITY_FISHING_BOBBER,
	EntitySpectralArrow: ENTITY_SPECTRAL_ARROW,
}

var Version = &minecraft.Version{
	Name:             "1.19",
	NameLatest:       "1.19",
	LoginClientCodec: LoginPacketClientCodec,
	LoginServerCodec: LoginPacketServerCodec,
	PlayClientCodec:  PlayPacketClientCodec,
	PlayServerCodec:  PlayPacketServerCodec,
	IdMap:            IdMap,
	Id: []int{
		759,
	},
}
