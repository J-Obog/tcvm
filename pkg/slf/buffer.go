package slf

import (
	"bytes"
	"encoding/binary"
)

func ReadU32(buf *bytes.Buffer) uint32 {
	data := buf.Next(4)
	return binary.BigEndian.Uint32(data)
}

func WriteU32(buf *bytes.Buffer, val uint32) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, val)
	buf.Write(b)
}

func ReadU8(buf *bytes.Buffer) uint8 {
	return buf.Next(1)[0]
}

func WriteU8(buf *bytes.Buffer, val uint8) {
	buf.Write([]byte{val})
}

func ReadStr(buf *bytes.Buffer, len uint32) string {
	str := buf.Next(int(len))
	return string(str)
}

func WriteStr(buf *bytes.Buffer, str string) {
	buf.Write([]byte(str))
}