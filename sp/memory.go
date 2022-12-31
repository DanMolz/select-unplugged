package sp

import (
	"errors"
	"fmt"
)

type Memory struct {
	area Area
	data Data
}

func NewMemory(area Area) Memory {
	memory := Memory{
		area: area,
	}
	return memory
}

func (m Memory) Area() Area {
	return m.area
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
	if m.data == nil {
		panic(errors.New("Read from uninitialized Memory"))
	}
	return m.data
}
