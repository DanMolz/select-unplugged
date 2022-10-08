package sp

import "encoding/binary"

type Address uint32

func (address Address) Message() Message {
	return binary.LittleEndian.AppendUint32(Message{}, uint32(address))
}
