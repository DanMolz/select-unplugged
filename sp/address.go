package sp

import (
	"encoding/binary"
	"errors"
)

type Address uint32

func (address Address) Message() Message {
	return binary.LittleEndian.AppendUint32(Message{}, uint32(address))
}

func NewAddressFromMessage(m Message) (*Address, error) {
	if len(m) < 6 {
		return nil, errors.New("Need more bytes to calculate address")
	}
	address := Address(binary.LittleEndian.Uint32(m[2:6]))
	return &address, nil
}
