package sp

import (
	"errors"
	"fmt"
)

type Memory struct {
	area Area
	data []byte
}

func (m *Memory) SetData(data []byte) error {
	if len(data) == m.area.Bytes() {
		m.data = data
		return nil
	}
	return errors.New(fmt.Sprintf("Got %d bytes, expecting %d", len(data), m.area.Bytes()))
}

func (m Memory) Data() []byte {
	return m.data
}
