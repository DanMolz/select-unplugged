package sp

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

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

// Calculate request length given enough bytes from a Message
func CalculateRequestLength(partial Message) (*int, error) {
	i := func(i int) *int { return &i }
	if len(partial) < 1 {
		return nil, errors.New("Need a byte to calculate length")
	}
	rt := MessageType(partial[0])
	length := 8
	if rt == Query {
		return i(8), nil
	}
	if rt != Write {
		return nil, errors.New("Unknown message type")
	}
	words, err := partial.Words()
	if err != nil {
		return nil, err
	}
	return i(length + *words*2 + 2), nil
}

func (r Request) ResponseLength() (*int, error) {
	requestType, err := r.Type()
	if err != nil {
		return nil, err
	}
	requestLength := len(r)
	dataLength := r.DataLength()
	if err != nil {
		return nil, err
	}
	crcLength := 2
	if *requestType == Write {
		return &requestLength, nil
	}
	length := requestLength + dataLength + crcLength
	return &length, nil
}

func (r Request) DataLength() int {
	return (int(r[1]) + 1) * 2
}

func (r Request) Type() (*MessageType, error) {
	return Message(r).Type()
}
