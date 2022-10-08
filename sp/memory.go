package sp

import (
	"bytes"
	"errors"
	"fmt"
)

type Memory struct {
	address Address
	data    Data
}

func NewMemory(address Address, words Words) (*Memory, error) {
	if words == 0 {
		return nil, errors.New("Non-zero word length required")
	}
	memory := Memory{
		address: address,
		data:    bytes.Repeat([]byte("\x00"), int(words)*2),
	}
	return &memory, nil

}

func (m *Memory) SetData(data Data) error {
	if len(data) == len(m.data) {
		m.data = data
		return nil
	}
	return errors.New(fmt.Sprintf("Got %d bytes, expecting %d", len(data), len(m.data)))
}

func (m Memory) Data() Data {
	return m.data
}
