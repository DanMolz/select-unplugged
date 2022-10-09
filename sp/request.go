package sp

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

type RequestType string

const Query RequestType = "Q"
const Write RequestType = "W"

type Request Message

func NewRequestQuery(area Area) Request {

	message := []byte("Q")
	message = append(message, area.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))

	return Request(message)
}

func NewRequestWrite(memory Memory) Request {
	message := Message("W")
	message = append(message, uint8(len(memory.data)/2-1))
	message = append(message, memory.address.Message()...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))
	message = append(message, memory.data...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))
	return Request(message)
}

func (r Request) String() string {
	return fmt.Sprintf("Request(0x%s)", hex.EncodeToString(r))
}

func (r Request) ResponseLength() (*int, error) {
	requestType, err := r.Type()
	if err != nil {
		return nil, err
	}
	requestLength := len(r)
	dataLength := (int(r[1]) + 1) * 2
	crcLength := 2
	if *requestType == Write {
		return &requestLength, nil
	}
	length := requestLength + dataLength + crcLength
	return &length, nil
}

func (r Request) Type() (*RequestType, error) {
	rt := RequestType(r[0])
	if rt == Query || rt == Write {
		return &rt, nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown request type '%s'", rt))
}
