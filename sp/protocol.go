package sp

import (
	"bytes"
	"errors"
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

func (protocol *Protocol) Login(password string) error {
	loginHashMemory := VarLoginHash.Memory()

	readLoginHashRequest := NewRequestQuery(VarLoginHash.Area())
	loginHashResponse, _ := protocol.Send(readLoginHashRequest)
	loginHash, err := Message(loginHashResponse).Data()
	if err != nil {
		return err
	}

	responseHash := CalculateLoginHash(password, *loginHash)

	loginHashMemory.SetData(Data(responseHash))
	writeLoginHashRequest := NewRequestWrite(loginHashMemory)
	_, err = protocol.Send(writeLoginHashRequest)
	if err != nil {
		return err
	}

	readLoginStatusRequest := NewRequestQuery(VarLoginStatus.Area())
	readLoginStatusResponse, err := protocol.Send(readLoginStatusRequest)
	if err != nil {
		return err
	}
	loginStatus, err := Message(readLoginStatusResponse).Data()
	if err != nil {
		return err
	}
	if !bytes.Equal(*loginStatus, []byte("\x01\x00")) {
		return errors.New("Invalid login status")
	}
	return nil
}
