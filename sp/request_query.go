package sp

import "encoding/binary"

type RequestQuery Request

func NewRequestQuery(area Area) RequestQuery {

	message := []byte("Q")
	message = append(message, area.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))

	return RequestQuery(message)
}
