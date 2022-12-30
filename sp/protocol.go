package sp

import (
	"sync"

	log "github.com/sirupsen/logrus"
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

func (protocol *Protocol) Send(request Request) (Response, error) {
	length, err := request.ResponseLength()
	if err != nil {
		return nil, err
	}
	protocol.mutex.Lock()
	defer protocol.mutex.Unlock()
	log.Debugf("> %s", request)
	protocol.connection.Write([]byte(request))
	data := make([]byte, 0)
	for len(data) < *length && err == nil {
		wantedLength := *length - len(data)
		buf := make([]byte, wantedLength)
		partlength, err := protocol.connection.Read(&buf)
		if err != nil {
			return nil, err
		}
		data = append(data, buf[:partlength]...)
	}
	log.Debugf("< %s", Response(data))
	return data, err
}
