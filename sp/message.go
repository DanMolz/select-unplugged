package sp

import (
	"encoding/hex"
	"fmt"
)

// Message is the wire level format of data within Request and Response.
type Message []byte

func (m Message) String() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(m))
}
