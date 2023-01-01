package sp

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
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
	return m.area.Words()
}

func (m Memory) Bytes() int {
	return int(m.Words() * 2)
}

func (m *Memory) SetData(data Data) {
	log.Debugf("Setting data to %x", data)
	if len(data) != m.Bytes() {
		panic(errors.New(fmt.Sprintf(
			"Got %d bytes, expecting %d",
			len(data),
			m.Bytes(),
		)))
	}
	m.data = data
}

func (m Memory) Data() Data {
	if m.data == nil {
		panic(errors.New("Read from uninitialized Memory"))
	}
	return m.data
}
