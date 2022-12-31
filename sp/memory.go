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

func NewMemory(address Address, words Words) Memory {
	if words == 0 {
		panic(errors.New("Non-zero word length required"))
	}
	memory := Memory{
		address: address,
		data:    bytes.Repeat([]byte("\x00"), int(words)*2),
	}
	return memory
}

func (m Memory) Address() Address {
	return m.address
}

func (m Memory) Words() Words {
	return Words(len(m.data) / 2)
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
