package minecraft

import (
	"bytes"
	"encoding/binary"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/klauspost/compress/zlib"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
)

type PacketGeneric struct {
	id         int
	Bytes      []byte
	compressed bool
	swappers   *PacketGenericSwappers
}

type PacketGenericSwappers struct {
	ClientInt           [][]int
	ClientVarInt        []bool
	ClientSpawnObject   int
	ClientEntityDestroy int
	ServerInt           [][]int
	ServerVarInt        []bool
	IdMap               *IdMap
}

const (
	PacketGenericSwappersClientSpawnObjectNoUUID int = iota
	PacketGenericSwappersClientSpawnObjectUUID
	PacketGenericSwappersClientSpawnObjectUUIDTypeVar
)

const (
	PacketGenericSwappersClientEntityDestroyVar int = iota
	PacketGenericSwappersClientEntityDestroyInt32
)

func NewPacketGeneric(id int, bytes []byte, compressed bool, swappers *PacketGenericSwappers) (this *PacketGeneric) {
	this = new(PacketGeneric)
	this.id = id
	this.Bytes = bytes
	this.compressed = compressed
	this.swappers = swappers
	return
}

func (this *PacketGeneric) Decompress() (err error) {
	if !this.compressed {
		return
	}
	buffer := bytes.NewReader(this.Bytes)
	_, err = packet.ReadVarInt(buffer) // compression length
	if err != nil {
		return
	}
	zlibReader, err := zlib.NewReader(buffer)
	if err != nil {
		return
	}
	_, err = packet.ReadVarInt(zlibReader) // id
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(zlibReader)
	if err != nil {
		return
	}
	this.Bytes = bytes
	this.compressed = false
	return
}

func (this *PacketGeneric) SwapEntities(a int32, b int32, isClientBound bool) {
	if a == b {
		return
	}
	if isClientBound {
		if this.id == this.swappers.IdMap.PacketClientSpawnObject {
			// PacketClientSpawnObject
			this.Decompress()
			buffer := bytes.NewReader(this.Bytes)
			// read
			entityId, _ := packet.ReadVarInt(buffer)
			var entityUUID uuid.UUID
			if this.swappers.ClientSpawnObject >= PacketGenericSwappersClientSpawnObjectUUID { // v1.9+ - has uuid
				entityUUID, _ = packet.ReadUUID(buffer)
			}
			var entityType int
			var entityTypeIsSwap bool
			if this.swappers.ClientSpawnObject >= PacketGenericSwappersClientSpawnObjectUUIDTypeVar { // v1.14+
				entityType, _ = packet.ReadVarInt(buffer)
				entityTypeIsSwap = entityType == this.swappers.IdMap.EntityArrow || entityType == this.swappers.IdMap.EntityFishingBobber || entityType == this.swappers.IdMap.EntitySpectralArrow
				if entityType == this.swappers.IdMap.EntityArrow || entityType == this.swappers.IdMap.EntitySpectralArrow {
					a = a + 1
					b = b + 1
				}
			} else {
				entityTypeU8, _ := packet.ReadUint8(buffer)
				entityType = int(entityTypeU8)
				entityTypeIsSwap = entityType == this.swappers.IdMap.EntityArrow || entityType == this.swappers.IdMap.EntityFishingBobber || entityType == this.swappers.IdMap.EntitySpectralArrow
				if (entityType == this.swappers.IdMap.EntityArrow || entityType == this.swappers.IdMap.EntitySpectralArrow) && this.swappers.ClientSpawnObject >= PacketGenericSwappersClientSpawnObjectUUID { // v1.9+
					a = a + 1
					b = b + 1
				}
			}
			if entityTypeIsSwap {
				var entitySkip []byte
				if this.swappers.ClientSpawnObject >= PacketGenericSwappersClientSpawnObjectUUID { // v1.9+ - f64 xyz
					entitySkip = make([]byte, 8+8+8+1+1)
				} else {
					entitySkip = make([]byte, 4+4+4+1+1)
				}
				buffer.Read(entitySkip)
				entityData, _ := packet.ReadInt32(buffer)
				// rewrite
				this.swapAndIf(entityData, a, b, func(newId int32) {
					entityData = newId
					bufferRewrite := new(bytes.Buffer)
					packet.WriteVarInt(bufferRewrite, entityId)
					if this.swappers.ClientSpawnObject >= PacketGenericSwappersClientSpawnObjectUUID { // v1.9+ - has uuid
						packet.WriteUUID(bufferRewrite, entityUUID)
					}
					if this.swappers.ClientSpawnObject >= PacketGenericSwappersClientSpawnObjectUUIDTypeVar { // v1.14+
						packet.WriteVarInt(bufferRewrite, entityType)
					} else {
						packet.WriteUint8(bufferRewrite, uint8(entityType))
					}
					bufferRewrite.Write(entitySkip)
					packet.WriteInt32(bufferRewrite, entityData)
					buffer.WriteTo(bufferRewrite)
					this.Bytes = bufferRewrite.Bytes()
				})
			}
		} else if this.id == this.swappers.IdMap.PacketClientCollectItem {
			// PacketClientCollectItem
			if ok := this.swappers.ClientVarInt[this.id]; ok {
				this.Decompress()
				buffer := bytes.NewReader(this.Bytes)
				collectedId, _ := packet.ReadVarInt(buffer)
				collectorId, _ := packet.ReadVarInt(buffer)
				this.swapAndIf(int32(collectorId), a, b, func(newId int32) {
					collectorId = int(newId)
					bufferRewrite := new(bytes.Buffer)
					packet.WriteVarInt(bufferRewrite, collectedId)
					packet.WriteVarInt(bufferRewrite, collectorId)
					buffer.WriteTo(bufferRewrite)
					this.Bytes = bufferRewrite.Bytes()
				})
			}
		} else if this.id == this.swappers.IdMap.PacketClientCombatEvent {
			// PacketClientCombatEvent
			this.Decompress()
			buffer := bytes.NewReader(this.Bytes)
			bufferRewrite := new(bytes.Buffer)

			eventId, _ := packet.ReadUint8(buffer)
			packet.WriteUint8(bufferRewrite, eventId)
			if eventId == 1 {
				// end combat
				duration, _ := packet.ReadVarInt(buffer)
				entityId, _ := packet.ReadInt32(buffer)
				entityId = this.swapAndRet(entityId, a, b)
				packet.WriteVarInt(bufferRewrite, duration)
				packet.WriteInt32(bufferRewrite, entityId)
			} else if eventId == 2 {
				// entity dead
				playerId, _ := packet.ReadVarInt(buffer)
				playerId = int(this.swapAndRet(int32(playerId), a, b))
				entityId, _ := packet.ReadInt32(buffer)
				entityId = this.swapAndRet(entityId, a, b)
				packet.WriteVarInt(bufferRewrite, playerId)
				packet.WriteInt32(bufferRewrite, entityId)
			}

			buffer.WriteTo(bufferRewrite)
			this.Bytes = bufferRewrite.Bytes()
		} else if this.id == this.swappers.IdMap.PacketClientEntitySoundEffect {
			// PacketClientEntitySoundEffect
			this.Decompress()
			buffer := bytes.NewReader(this.Bytes)
			soundId, _ := packet.ReadVarInt(buffer)
			soundCategory, _ := packet.ReadVarInt(buffer)
			entityId, _ := packet.ReadVarInt(buffer)
			this.swapAndIf(int32(entityId), a, b, func(newId int32) {
				entityId = int(newId)
				bufferRewrite := new(bytes.Buffer)
				packet.WriteVarInt(bufferRewrite, soundId)
				packet.WriteVarInt(bufferRewrite, soundCategory)
				packet.WriteVarInt(bufferRewrite, entityId)
				buffer.WriteTo(bufferRewrite)
				this.Bytes = bufferRewrite.Bytes()
			})
		} else if this.id == this.swappers.IdMap.PacketClientSetPassengers || this.id == this.swappers.IdMap.PacketClientDestroyEntities {
			// PacketClientSetPassengers & PacketClientDestroyEntities
			this.Decompress()
			buffer := bytes.NewReader(this.Bytes)
			bufferRewrite := new(bytes.Buffer)
			if this.id == this.swappers.IdMap.PacketClientSetPassengers {
				entityId, _ := packet.ReadVarInt(buffer)
				packet.WriteVarInt(bufferRewrite, entityId)
			}
			if this.swappers.ClientEntityDestroy == PacketGenericSwappersClientEntityDestroyInt32 { // 1.17
				nEntities, _ := packet.ReadUint8(buffer)
				packet.WriteUint8(bufferRewrite, nEntities)
				for i := uint8(0); i < nEntities; i++ {
					entityId, _ := packet.ReadInt32(buffer)
					entityId = this.swapAndRet(entityId, a, b)
					packet.WriteInt32(bufferRewrite, entityId)
				}
			} else { // 1.18+
				nEntities, _ := packet.ReadVarInt(buffer)
				packet.WriteVarInt(bufferRewrite, nEntities)
				for i := 0; i < nEntities; i++ {
					entityId, _ := packet.ReadVarInt(buffer)
					entityId = int(this.swapAndRet(int32(entityId), a, b))
					packet.WriteVarInt(bufferRewrite, entityId)
				}
			}
			this.Bytes = bufferRewrite.Bytes()
		}
	}
	// TODO entity metadata?
	this.swapEntitiesInt(a, b, isClientBound)
	this.swapEntitiesVarInt(a, b, isClientBound)
}

func (this *PacketGeneric) swapAndIf(field int32, a int32, b int32, f func(int32)) {
	if field == a {
		f(b)
	} else if field == b {
		f(a)
	}
}

func (this *PacketGeneric) swapAndRet(field int32, a int32, b int32) int32 {
	if field == a {
		return b
	} else if field == b {
		return a
	} else {
		return field
	}
}

func (this *PacketGeneric) swapEntitiesInt(a int32, b int32, isClientBound bool) {
	var positions [][]int
	if isClientBound {
		positions = this.swappers.ClientInt
	} else {
		positions = this.swappers.ServerInt
	}
	if this.id < 0 {
		return
	}
	if this.id >= len(positions) {
		return
	}
	idPositions := positions[this.id]
	if idPositions == nil {
		return
	}
	this.Decompress()
	var id int32
	for _, position := range idPositions {
		if len(this.Bytes) < position+4 {
			continue
		}
		id = int32(binary.BigEndian.Uint32(this.Bytes[position : position+4]))
		if id == a {
			binary.BigEndian.PutUint32(this.Bytes[position:position+4], uint32(b))
		} else if id == b {
			binary.BigEndian.PutUint32(this.Bytes[position:position+4], uint32(a))
		}
	}
}

func (this *PacketGeneric) swapEntitiesVarInt(a int32, b int32, isClientBound bool) {
	var positions []bool
	if isClientBound {
		positions = this.swappers.ClientVarInt
	} else {
		positions = this.swappers.ServerVarInt
	}
	if this.id < 0 {
		return
	}
	if this.id >= len(positions) {
		return
	}
	if positions[this.id] == false {
		return
	}
	this.Decompress()
	// Read the old Id
	buffer := bytes.NewBuffer(this.Bytes)
	id, err := packet.ReadVarInt(buffer)
	if err != nil {
		return
	}
	// Check the Id
	var newId int
	if id == int(a) {
		newId = int(b)
	} else if id == int(b) {
		newId = int(a)
	} else {
		return
	}
	// Apply the new Id
	bufferRewrite := new(bytes.Buffer)
	err = packet.WriteVarInt(bufferRewrite, newId)
	if err != nil {
		return
	}
	buffer.WriteTo(bufferRewrite)
	this.Bytes = bufferRewrite.Bytes()
}

func (this *PacketGeneric) Raw() bool {
	return this.compressed
}

func (this *PacketGeneric) Id() int {
	return this.id
}

type packetGenericCodec struct {
	Id       int
	swappers *PacketGenericSwappers
}

func NewPacketGenericCodec(id int, swappers *PacketGenericSwappers) (this *packetGenericCodec) {
	this = new(packetGenericCodec)
	this.Id = id
	this.swappers = swappers
	return
}

func (this *packetGenericCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetGeneric := new(PacketGeneric)
	packetGeneric.id = this.Id
	packetGeneric.swappers = this.swappers
	if zlibReader, ok := reader.(*packet.ZlibToggleReader); ok {
		zlibReader.SetRaw(true)
		packetGeneric.compressed = true
	}
	packetGeneric.Bytes, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	decode = packetGeneric
	return
}

func (this *packetGenericCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	_, err = writer.Write(encode.(*PacketGeneric).Bytes)
	return
}
