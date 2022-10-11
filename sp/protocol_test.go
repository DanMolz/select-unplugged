package sp

import (
	"testing"

	"github.com/neerolyte/select-unplugged/sp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProtocolSend(t *testing.T) {
	connection := mocks.NewConnection(t)

	req := NewRequestQuery(Area{
		address: 0x01,
		words:   1,
	})
	connection.On("Write", []byte(req)).Return()
	connection.On("Read", 14).Return([]byte("some-response"))

	protocol := NewProtocol(connection)
	res, err := protocol.Send(req)

	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("some-response"), []byte(res))
}
