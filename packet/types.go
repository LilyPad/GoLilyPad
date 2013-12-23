package packet

import "encoding/binary"
import "errors"
import "fmt"
import "io"

func WriteString(writer io.Writer, util []byte, s string) (err error) {
	bytes := []byte(s)
	err = WriteVarInt(writer, util, len(bytes))
	if err != nil {
		return
	}
	_, err = writer.Write(bytes)
	return
}

func ReadString(reader io.Reader, util []byte) (s string, err error) {
	length, err := ReadVarInt(reader, util)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("String length is below zero: %d", length))
		return
	}
	if length > 2097151 { // 2^21
		err = errors.New(fmt.Sprintf("String length is above maximum: %d", length))
		return
	}
	bytes := make([]byte, length)
	_, err = reader.Read(bytes)
	if err != nil {
		return
	}
	return string(bytes), nil
}

func WriteVarInt(writer io.Writer, util []byte, val int) (err error) {
	for val >= 0x80 {
		err = WriteUint8(writer, util, byte(val) | 0x80)
		if err != nil {
			return
		}
		val >>= 7
	}
	return WriteUint8(writer, util, byte(val))
}

func ReadVarInt(reader io.Reader, util []byte) (val int, err error) {
	var s uint
	var b byte
	i := 0
	for {
		b, err = ReadUint8(reader, util)
		if err != nil {
			return
		}
		if b < 0x80 {
			return (val | int(b) << s), nil
		}
		if i > 5 {
			return 0, errors.New("VarInt too long")
		}
		val |= int(b & 0x7f) << s
		s += 7
		i++
	}
	return 0, nil
}

func ReadBool(reader io.Reader, util []byte) (val bool, err error) {
	uval, err := ReadUint8(reader, util)
	if err != nil {
		return
	}
	return uval != 0, nil
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
	return int8(uval), nil
}

func WriteInt8(writer io.Writer, util []byte, val int8) (err error) {
	return WriteUint8(writer, util, uint8(val))
}

func ReadUint8(reader io.Reader, util []byte) (val uint8, err error) {
	_, err = reader.Read(util[:1])
	if err != nil {
		return
	}
	return util[0], nil
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
	return int16(uval), nil
}

func WriteInt16(writer io.Writer, util []byte, val int16) (err error) {
	return WriteUint16(writer, util, uint16(val))
}

func ReadUint16(reader io.Reader, util []byte) (val uint16, err error) {
	_, err = reader.Read(util[:2])
	if err != nil {
		return
	}
	return binary.BigEndian.Uint16(util[:2]), nil
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
	return int32(uval), nil
}

func WriteInt32(writer io.Writer, util []byte, val int32) (err error) {
	return WriteUint32(writer, util, uint32(val))
}

func ReadUint32(reader io.Reader, util []byte) (val uint32, err error) {
	_, err = reader.Read(util[:4])
	if err != nil {
		return
	}
	return binary.BigEndian.Uint32(util[:4]), nil
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
	return int64(uval), nil
}

func WriteInt64(writer io.Writer, util []byte, val int64) (err error) {
	return WriteUint64(writer, util, uint64(val))
}

func ReadUint64(reader io.Reader, util []byte) (val uint64, err error) {
	_, err = reader.Read(util[:8])
	if err != nil {
		return
	}
	return binary.BigEndian.Uint64(util[:8]), nil
}

func WriteUint64(writer io.Writer, util []byte, val uint64) (err error) {
	binary.BigEndian.PutUint64(util[:8], val)
	_, err = writer.Write(util[:8])
	return
}
