package sp

import "encoding/binary"

type RequestWrite Request

func NewRequestWrite(memory Memory) RequestWrite {
	message := Message("W")
	message = append(message, uint8(len(memory.data)/2-1))
	message = append(message, memory.address.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))
	message = append(message, memory.data...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))
	return RequestWrite(message)
}
