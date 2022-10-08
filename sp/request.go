package sp

import (
	"encoding/hex"
	"fmt"
)

type Request Message

func (r Request) String() string {
	return fmt.Sprintf("Request(0x%s)", hex.EncodeToString(r))
}
