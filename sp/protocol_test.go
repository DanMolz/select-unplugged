package sp

import (
	"testing"
)

type ConnectionMock struct {
	Connection
}

func (ConnectionMock) Read(length int) []byte {
	return []byte("raw")
}

func (ConnectionMock) Write(data []byte) {
}

func TestProtocolLocking(t *testing.T) {
	t.Fatal("TODO: test locking")
	connection := ConnectionMock{}

	protocol := NewProtocol(connection)
	protocol.Send(Request(NewRequestQuery(Area{
		address: 0x00,
		words:   1,
	})))
}
