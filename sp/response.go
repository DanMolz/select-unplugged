package sp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Response message from SP Pro.
type Response Message

func (r Response) String() string {
	return fmt.Sprintf("Response(0x%s)", hex.EncodeToString(r))
}

func NewResponse(request Request, data Message) (Response, error) {
	response := request
	response = append(request, data...)
	response = binary.LittleEndian.AppendUint16(response, Crc(response))
	return Response(response), nil
}
