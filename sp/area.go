// memory area
package sp

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type Area struct {
	address Address
	words   Words
}

func NewArea(address Address, words Words) Area {
	if words == 0 {
		panic(errors.New("Non-zero word length required"))
	}
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

func (area Area) Address() Address {
	return area.address
}

func (area Area) Words() Words {
	return area.words
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
