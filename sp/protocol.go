package sp

import (
	"sync"
)

// Protocol sends requests and receives responses via the connection.
type Protocol struct {
	connection *Connection
	mutex      sync.Mutex
}

func NewProtocol(connection *Connection) *Protocol {
	protocol := Protocol{
		connection: connection,
	}
	return &protocol
}

func (protocol *Protocol) Send(request Request) Message {
	protocol.mutex.Lock()
	defer protocol.mutex.Unlock()
	protocol.connection.write(request)
	return protocol.connection.read(1024) // TODO: get read length from request
}
