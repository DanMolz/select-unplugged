package sp

import (
	"encoding/binary"
)

type Request []byte

func CreateQueryRequest(area Area) Request {

	message := []byte("Q")
	message = append(message, area.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))

	return Request(message)
}
