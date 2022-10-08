package sp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Request []byte

func (r Request) String() string {
	return fmt.Sprintf("Request(0x%s)", hex.EncodeToString(r))
}

func CreateQueryRequest(area Area) Request {

	message := []byte("Q")
	message = append(message, area.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))

	return Request(message)
}

func CreateWriteRequest(memory Memory) Request {

	message := []byte("W")
	message = append(
		message,
		uint8(len(memory.data)/2-1),
	)

	message = append(message, memory.address.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))
	message = append(message, memory.data...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))

	return Request(message)
}
