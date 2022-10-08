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

/*func CreateWriteRequest(memory Memory) Request {

	message := []byte("W")
	message = append(
		message,
		[]byte{uint8(len(memory.data))},
	)

	message = append(message, memory.area.Message()...)
	message = append(message, memory.data...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))

	return Request(message)
}
*/
