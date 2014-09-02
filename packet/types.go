package packet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	uuid "code.google.com/p/go-uuid/uuid"
)

func WriteString(writer io.Writer, util []byte, val string) (err error) {
	bytes := []byte(val)
	err = WriteVarInt(writer, util, len(bytes))
	if err != nil {
		return
	}
	_, err = writer.Write(bytes)
	return
}

func ReadString(reader io.Reader, util []byte) (val string, err error) {
	length, err := ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Decode, String length is below zero: %d", length))
		return
	}
	if length > 1048576 { // 2^(21-1)
		err = errors.New(fmt.Sprintf("Decode, String length is above maximum: %d", length))
		return
	}
	bytes := make([]byte, length)
	_, err = reader.Read(bytes)
	if err != nil {
		return
	}
	val = string(bytes)
	return
}

func WriteVarInt(writer io.Writer, util []byte, val int) (err error) {
	for val >= 0x80 {
		err = WriteUint8(writer, util, byte(val) | 0x80)
		if err != nil {
			return
		}
		val >>= 7
	}
	err = WriteUint8(writer, util, byte(val))
	return
}

func ReadVarInt(reader io.Reader, util []byte) (result int, err error) {
	var bytes byte = 0
	var b byte
	for {
		b, err = ReadUint8(reader, util)
		if err != nil {
			return
		}
		result |= int(uint(b & 0x7F) << uint(bytes * 7))
		bytes++
		if bytes > 5 {
			err = errors.New("Decode, VarInt is too long")
			return
		}
		if (b & 0x80) == 0x80 {
			continue
		}
		break
	}
	return
}

func WriteUUID(writer io.Writer, util []byte, val uuid.UUID) (err error) {
	_, err = writer.Write(val)
	return
}

func ReadUUID(reader io.Reader, util []byte) (result uuid.UUID, err error) {
	bytes := make([]byte, 16)
	_, err = reader.Read(bytes)
	if err != nil {
		return
	}
	result = uuid.UUID(bytes)
	return
}

func ReadBool(reader io.Reader, util []byte) (val bool, err error) {
	uval, err := ReadUint8(reader, util)
	if err != nil {
		return
	}
	val = uval != 0
	return
}

func WriteBool(writer io.Writer, util []byte, val bool) (err error) {
	if val {
		err = WriteUint8(writer, util, 1)
		return
	}
	err = WriteUint8(writer, util, 0)
	return
}

func ReadInt8(reader io.Reader, util []byte) (val int8, err error) {
	uval, err := ReadUint8(reader, util)
	if err != nil {
		return
	}
	val = int8(uval)
	return
}

func WriteInt8(writer io.Writer, util []byte, val int8) (err error) {
	err = WriteUint8(writer, util, uint8(val))
	return
}

func ReadUint8(reader io.Reader, util []byte) (val uint8, err error) {
	_, err = reader.Read(util[:1])
	if err != nil {
		return
	}
	val = util[0]
	return
}

func WriteUint8(writer io.Writer, util []byte, val uint8) (err error) {
	util[0] = val
	_, err = writer.Write(util[:1])
	return
}

func ReadInt16(reader io.Reader, util []byte) (val int16, err error) {
	uval, err := ReadUint16(reader, util)
	if err != nil {
		return
	}
	val = int16(uval)
	return
}

func WriteInt16(writer io.Writer, util []byte, val int16) (err error) {
	err = WriteUint16(writer, util, uint16(val))
	return
}

func ReadUint16(reader io.Reader, util []byte) (val uint16, err error) {
	_, err = reader.Read(util[:2])
	if err != nil {
		return
	}
	val = binary.BigEndian.Uint16(util[:2])
	return
}

func WriteUint16(writer io.Writer, util []byte, val uint16) (err error) {
	binary.BigEndian.PutUint16(util[:2], val)
	_, err = writer.Write(util[:2])
	return
}

func ReadInt32(reader io.Reader, util []byte) (val int32, err error) {
	uval, err := ReadUint32(reader, util)
	if err != nil {
		return
	}
	val = int32(uval)
	return
}

func WriteInt32(writer io.Writer, util []byte, val int32) (err error) {
	err = WriteUint32(writer, util, uint32(val))
	return
}

func ReadUint32(reader io.Reader, util []byte) (val uint32, err error) {
	_, err = reader.Read(util[:4])
	if err != nil {
		return
	}
	val = binary.BigEndian.Uint32(util[:4])
	return
}

func WriteUint32(writer io.Writer, util []byte, val uint32) (err error) {
	binary.BigEndian.PutUint32(util[:4], val)
	_, err = writer.Write(util[:4])
	return
}

func ReadInt64(reader io.Reader, util []byte) (val int64, err error) {
	uval, err := ReadUint64(reader, util)
	if err != nil {
		return
	}
	val = int64(uval)
	return
}

func WriteInt64(writer io.Writer, util []byte, val int64) (err error) {
	err = WriteUint64(writer, util, uint64(val))
	return
}

func ReadUint64(reader io.Reader, util []byte) (val uint64, err error) {
	_, err = reader.Read(util[:8])
	if err != nil {
		return
	}
	val = binary.BigEndian.Uint64(util[:8])
	return
}

func WriteUint64(writer io.Writer, util []byte, val uint64) (err error) {
	binary.BigEndian.PutUint64(util[:8], val)
	_, err = writer.Write(util[:8])
	return
}
