package sp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Response message from SP Pro.
type Response struct {
	message Message
}

func (r Response) String() string {
	return fmt.Sprintf("Response(0x%s)", hex.EncodeToString(r.Message()))
}

func (r Response) Message() Message {
	return r.message
}

func NewResponse(request Request, data Message) (Response, error) {
	message := request.Message()
	message = append(message, data...)
	message = binary.LittleEndian.AppendUint16(message, Crc(message))
	return Response{
		message: message,
	}, nil
}
