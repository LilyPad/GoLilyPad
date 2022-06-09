package packet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/auth"
	uuid "github.com/satori/go.uuid"
	"io"
	"math"
)

func WriteString(writer io.Writer, val string) (err error) {
	bytes := []byte(val)
	err = WriteVarInt(writer, len(bytes))
	if err != nil {
		return
	}
	_, err = writer.Write(bytes)
	return
}

func ReadString(reader io.Reader) (val string, err error) {
	length, err := ReadVarInt(reader)
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

func WriteVarInt(writer io.Writer, val int) (err error) {
	uval := uint32(val)
	for uval >= 0x80 {
		err = WriteUint8(writer, byte(uval)|0x80)
		if err != nil {
			return
		}
		uval >>= 7
	}
	err = WriteUint8(writer, byte(uval))
	return
}

func ReadVarInt(reader io.Reader) (result int, err error) {
	var bytes byte = 0
	var b byte
	var uresult uint32 = 0
	for {
		b, err = ReadUint8(reader)
		if err != nil {
			return
		}
		uresult |= uint32(b&0x7F) << uint32(bytes*7)
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
	result = int(int32(uresult))
	return
}

func WriteUUID(writer io.Writer, val uuid.UUID) (err error) {
	_, err = writer.Write(val[:])
	return
}

func ReadUUID(reader io.Reader) (result uuid.UUID, err error) {
	bytes := make([]byte, 16)
	_, err = reader.Read(bytes)
	if err != nil {
		return
	}
	result, _ = uuid.FromBytes(bytes)
	return
}

func ReadBool(reader io.Reader) (val bool, err error) {
	uval, err := ReadUint8(reader)
	if err != nil {
		return
	}
	val = uval != 0
	return
}

func WriteBool(writer io.Writer, val bool) (err error) {
	if val {
		err = WriteUint8(writer, 1)
	} else {
		err = WriteUint8(writer, 0)
	}
	return
}

func ReadInt8(reader io.Reader) (val int8, err error) {
	uval, err := ReadUint8(reader)
	val = int8(uval)
	return
}

func WriteInt8(writer io.Writer, val int8) (err error) {
	err = WriteUint8(writer, uint8(val))
	return
}

func ReadUint8(reader io.Reader) (val uint8, err error) {
	var util [1]byte
	_, err = reader.Read(util[:1])
	val = util[0]
	return
}

func WriteUint8(writer io.Writer, val uint8) (err error) {
	var util [1]byte
	util[0] = val
	_, err = writer.Write(util[:1])
	return
}

func ReadInt16(reader io.Reader) (val int16, err error) {
	uval, err := ReadUint16(reader)
	val = int16(uval)
	return
}

func WriteInt16(writer io.Writer, val int16) (err error) {
	err = WriteUint16(writer, uint16(val))
	return
}

func ReadUint16(reader io.Reader) (val uint16, err error) {
	var util [2]byte
	_, err = reader.Read(util[:2])
	val = binary.BigEndian.Uint16(util[:2])
	return
}

func WriteUint16(writer io.Writer, val uint16) (err error) {
	var util [2]byte
	binary.BigEndian.PutUint16(util[:2], val)
	_, err = writer.Write(util[:2])
	return
}

func ReadInt32(reader io.Reader) (val int32, err error) {
	uval, err := ReadUint32(reader)
	val = int32(uval)
	return
}

func WriteInt32(writer io.Writer, val int32) (err error) {
	err = WriteUint32(writer, uint32(val))
	return
}

func ReadUint32(reader io.Reader) (val uint32, err error) {
	var util [4]byte
	_, err = reader.Read(util[:4])
	val = binary.BigEndian.Uint32(util[:4])
	return
}

func WriteUint32(writer io.Writer, val uint32) (err error) {
	var util [4]byte
	binary.BigEndian.PutUint32(util[:4], val)
	_, err = writer.Write(util[:4])
	return
}

func ReadInt64(reader io.Reader) (val int64, err error) {
	uval, err := ReadUint64(reader)
	val = int64(uval)
	return
}

func WriteInt64(writer io.Writer, val int64) (err error) {
	err = WriteUint64(writer, uint64(val))
	return
}

func ReadUint64(reader io.Reader) (val uint64, err error) {
	var util [8]byte
	_, err = reader.Read(util[:8])
	val = binary.BigEndian.Uint64(util[:8])
	return
}

func WriteUint64(writer io.Writer, val uint64) (err error) {
	var util [8]byte
	binary.BigEndian.PutUint64(util[:8], val)
	_, err = writer.Write(util[:8])
	return
}

func WriteFloat32(writer io.Writer, val float32) (err error) {
	return WriteUint32(writer, math.Float32bits(val))
}

func ReadFloat32(reader io.Reader) (val float32, err error) {
	ival, err := ReadUint32(reader)
	val = math.Float32frombits(ival)
	return
}

func WriteFloat64(writer io.Writer, val float64) (err error) {
	return WriteUint64(writer, math.Float64bits(val))
}

func ReadFloat64(reader io.Reader) (val float64, err error) {
	ival, err := ReadUint64(reader)
	val = math.Float64frombits(ival)
	return
}

func WriteArrByte(writer io.Writer, value []byte) error {
	err := WriteVarInt(writer, len(value))
	if err != nil {
		return err
	}
	_, err = writer.Write(value)
	return err
}

func ReadArrByte(reader io.Reader) ([]byte, error) {
	return ReadArrByteLimit(reader, -1)
}

func ReadArrByteLimit(reader io.Reader, maxLength int) ([]byte, error) {
	length, err := ReadVarInt(reader)
	if err != nil {
		return nil, err
	}
	if maxLength != -1 && length > maxLength {
		return nil, fmt.Errorf("array length %d is greater than max length %d", length, maxLength)
	}

	result := make([]byte, length)
	_, err = reader.Read(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ReadArrGameProfileProperty(reader io.Reader) ([]auth.GameProfileProperty, error) {
	length, err := ReadVarInt(reader)
	if err != nil {
		return nil, err
	}
	result := make([]auth.GameProfileProperty, length)
	for i := 0; i < length; i++ {
		result[i], err = ReadGameProfileProperty(reader)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func WriteArrGameProfileProperty(writer io.Writer, properties []auth.GameProfileProperty) error {
	err := WriteVarInt(writer, len(properties))
	if err != nil {
		return err
	}
	for _, property := range properties {
		err = WriteGameProfileProperty(writer, property)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadGameProfileProperty(reader io.Reader) (result auth.GameProfileProperty, err error) {
	result = auth.GameProfileProperty{}
	result.Name, err = ReadString(reader)
	if err != nil {
		return result, err
	}
	result.Value, err = ReadString(reader)
	if err != nil {
		return result, err
	}
	hasSignature, err := ReadBool(reader)
	if err != nil {
		return result, err
	}
	if hasSignature {
		result.Signature, err = ReadString(reader)
		if err != nil {
			return result, err
		}
	} else {
		result.Signature = ""
	}
	return result, nil
}

func WriteGameProfileProperty(writer io.Writer, property auth.GameProfileProperty) error {
	err := WriteString(writer, property.Name)
	if err != nil {
		return err
	}
	err = WriteString(writer, property.Value)
	if err != nil {
		return err
	}
	err = WriteBool(writer, property.HasSignature())
	if err != nil {
		return err
	}
	if property.HasSignature() {
		return WriteString(writer, property.Signature)
	}
	return nil
}
