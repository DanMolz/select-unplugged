// memory area
package sp

import (
	"encoding/binary"
	"fmt"
)

type Area struct {
	address Address
	words   Words
}

func NewArea(address Address, words Words) Area {
	return Area{
		address: address,
		words:   words,
	}
}

func (area Area) String() string {
	return fmt.Sprintf(
		"Area(0x%08x, %v)",
		area.address,
		area.words,
	)
}

func (area Area) Message() Message {
	return append(
		[]byte{uint8(area.words - 1)},
		binary.LittleEndian.AppendUint32(Message{}, uint32(area.address))...,
	)
}

func (area Area) Bytes() int {
	return int(area.words) * 2
}
