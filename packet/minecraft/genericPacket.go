package minecraft

import (
	"bytes"
	"encoding/binary"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/klauspost/compress/zlib"
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
	ClientInt    [][]int
	ClientVarInt []bool
	ServerInt    [][]int
	ServerVarInt []bool
	IdMap        *IdMap
}

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

func (this *PacketGeneric) SwapEntities(a int32, b int32, clientServer bool) {
	if a == b {
		return
	}
	newBuffer := &bytes.Buffer{}
	if this.id == this.swappers.IdMap.PacketClientSpawnObject && clientServer {
		this.Decompress()
		buffer := bytes.NewReader(this.Bytes)
		if this.id == 0x00 { // v1.9
			entityId, _ := packet.ReadVarInt(buffer)
			entityUUID, _ := packet.ReadUUID(buffer)
			entityType, _ := packet.ReadUint8(buffer)
			entityX, _ := packet.ReadFloat64(buffer)
			entityY, _ := packet.ReadFloat64(buffer)
			entityZ, _ := packet.ReadFloat64(buffer)
			entityPitch, _ := packet.ReadUint8(buffer)
			entityYaw, _ := packet.ReadUint8(buffer)
			entityData, _ := packet.ReadInt32(buffer)
			if (entityType >= 60 && entityType <= 66) || entityType == 90 {
				if entityData == a {
					entityData = b
				} else if entityData == b {
					entityData = a
				}
			}
			packet.WriteVarInt(newBuffer, entityId)
			packet.WriteUUID(newBuffer, entityUUID)
			packet.WriteUint8(newBuffer, entityType)
			packet.WriteFloat64(newBuffer, entityX)
			packet.WriteFloat64(newBuffer, entityY)
			packet.WriteFloat64(newBuffer, entityZ)
			packet.WriteUint8(newBuffer, entityPitch)
			packet.WriteUint8(newBuffer, entityYaw)
			packet.WriteInt32(newBuffer, entityData)
		} else {
			entityId, _ := packet.ReadVarInt(buffer)
			entityType, _ := packet.ReadUint8(buffer)
			entityX, _ := packet.ReadInt32(buffer)
			entityY, _ := packet.ReadInt32(buffer)
			entityZ, _ := packet.ReadInt32(buffer)
			entityPitch, _ := packet.ReadUint8(buffer)
			entityYaw, _ := packet.ReadUint8(buffer)
			entityData, _ := packet.ReadInt32(buffer)
			if (entityType >= 60 && entityType <= 66) || entityType == 90 {
				if entityData == a {
					entityData = b
				} else if entityData == b {
					entityData = a
				}
			}
			packet.WriteVarInt(newBuffer, entityId)
			packet.WriteUint8(newBuffer, entityType)
			packet.WriteInt32(newBuffer, entityX)
			packet.WriteInt32(newBuffer, entityY)
			packet.WriteInt32(newBuffer, entityZ)
			packet.WriteUint8(newBuffer, entityPitch)
			packet.WriteUint8(newBuffer, entityYaw)
			packet.WriteInt32(newBuffer, entityData)
		}
		buffer.WriteTo(newBuffer)
		this.Bytes = newBuffer.Bytes()
	} else if (this.id == this.swappers.IdMap.PacketClientSetPassengers /* || this.id == this.swappers.IdMap.PacketClientDestroyEntities*/) && clientServer {
		this.Decompress()
		buffer := bytes.NewReader(this.Bytes)
		if this.id == this.swappers.IdMap.PacketClientSetPassengers {
			entityId, _ := packet.ReadVarInt(buffer)
			packet.WriteVarInt(newBuffer, entityId)
		}
		nEntities, _ := packet.ReadVarInt(buffer)
		packet.WriteVarInt(newBuffer, nEntities)
		for i := 0; i < nEntities; i++ {
			entityId, _ := packet.ReadVarInt(buffer)
			if entityId == int(a) {
				entityId = int(b)
			} else if entityId == int(b) {
				entityId = int(a)
			}
			packet.WriteVarInt(newBuffer, entityId)
		}
		this.Bytes = newBuffer.Bytes()
	}
	// FIXME combat event
	// FIXME collect item
	this.swapEntitiesInt(a, b, clientServer)
	this.swapEntitiesVarInt(a, b, clientServer)
}

func (this *PacketGeneric) swapEntitiesInt(a int32, b int32, clientServer bool) {
	var positions [][]int
	if clientServer {
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

func (this *PacketGeneric) swapEntitiesVarInt(a int32, b int32, clientServer bool) {
	var positions []bool
	if clientServer {
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
	newBuffer := new(bytes.Buffer)
	err = packet.WriteVarInt(newBuffer, newId)
	if err != nil {
		return
	}
	buffer.WriteTo(newBuffer)
	this.Bytes = newBuffer.Bytes()
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
