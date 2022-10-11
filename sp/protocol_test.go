package sp

import (
	"testing"

	"github.com/neerolyte/select-unplugged/sp/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProtocolSend(t *testing.T) {
	connection := mocks.NewConnection(t)

	req := NewRequestQuery(Area{
		address: 0x01,
		words:   1,
	})
	connection.On("Write", []byte(req)).Return(len(req), nil)
	buf := make([]byte, 14)
	connection.On("Read", &buf).Return(14, nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]byte)
		*arg = []byte(Message("some-response"))
	})

	protocol := NewProtocol(connection)
	res, err := protocol.Send(req)

	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("some-response"), []byte(res))
}
