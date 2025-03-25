package bson

import "encoding/binary"

func GetBsonBytesLength(lengthSlice []byte) int {
	return int(binary.LittleEndian.Uint32(lengthSlice))
}
