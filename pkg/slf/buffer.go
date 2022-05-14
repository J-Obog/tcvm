package slf

import (
	"bytes"
	"encoding/binary"
)

func ReadU32(buf *bytes.Buffer) uint32 {
	data := buf.Next(4)
	return binary.BigEndian.Uint32(data)
}

func ReadU8(buf *bytes.Buffer) uint8 {
	return buf.Next(1)[0]
}

func ReadStr(buf *bytes.Buffer, len uint32) string {
	str := buf.Next(int(len))
	return string(str)
}