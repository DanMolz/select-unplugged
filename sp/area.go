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

func (area Area) String() string {
	return fmt.Sprintf(
		"Area(0x%08x, %v)",
		area.address,
		area.words,
	)
}

func (area Area) Message() Message {
	return append(
		[]byte{uint8(area.words)},
		binary.LittleEndian.AppendUint32(Message{}, uint32(area.address))...,
	)
}

func (area Area) Bytes() int {
	return int(area.words) * 2
}
