package sp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProtocolSend(t *testing.T) {
	connection := NewMockConnection(t)

	req := NewRequestQuery(Area{
		address: 0x01,
		words:   2,
	})
	connection.On("Write", []byte(req)).Return(len(req), nil)
	buf := make([]byte, 14)
	connection.On("Read", &buf).Return(14, nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]byte)
		*arg = []byte(Message([]byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d")))
	})

	protocol := NewProtocol(connection)
	res, err := protocol.Send(req)

	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d"), []byte(res))
	assert.Equal(t, 14, len(res))
}
