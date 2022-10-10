package sp

import (
	"sync"
)

// Protocol sends requests and receives responses via the connection.
type Protocol struct {
	connection Connection
	mutex      sync.Mutex
}

func NewProtocol(connection Connection) *Protocol {
	protocol := Protocol{
		connection: connection,
	}
	return &protocol
}

func (protocol *Protocol) Send(request Request) (Message, error) {
	length, err := request.ResponseLength()
	if err != nil {
		return nil, err
	}
	protocol.mutex.Lock()
	defer protocol.mutex.Unlock()
	protocol.connection.Write(request)
	return protocol.connection.Read(*length), nil
}
