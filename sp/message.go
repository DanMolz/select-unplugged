package sp

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type MessageType string

const Query MessageType = "Q"
const Write MessageType = "W"

// Message is the wire level format of data within Request and Response.
type Message []byte

func (m Message) String() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(m))
}

func (m Message) Describe() string {
	messageType, err := m.Type()
	if err != nil {
		panic(err)
	}
	messageAddress, err := m.Address()
	if err != nil {
		panic(err)
	}
	messageData, err := m.Data()
	dataDescription := ""
	if messageData != nil {
		dataDescription = fmt.Sprintf("=0x%x", *messageData)
	}
	return fmt.Sprintf("%s@%d%s", *messageType, *messageAddress, dataDescription)
}

func (m Message) Type() (*MessageType, error) {
	mt := MessageType(m[0])
	if mt != Query && mt != Write {
		return nil, errors.New(fmt.Sprintf("Unknown message type '%s'", mt))
	}
	return &mt, nil
}

func (m Message) Words() (*int, error) {
	if len(m) < 2 {
		return nil, errors.New("Need more bytes to calculate length")
	}
	words := int(m[1]) + 1
	return &words, nil
}

func (m Message) Address() (*Address, error) {
	return NewAddressFromMessage(m)
}

func (m Message) Data() (*[]byte, error) {
	words, err := m.Words()
	if err != nil {
		return nil, err
	}
	last := 8 + (*words * 2)
	if len(m) < last {
		return nil, errors.New("Need more bytes to return data")
	}
	data := []byte(m[8:last])
	return &data, nil
}
