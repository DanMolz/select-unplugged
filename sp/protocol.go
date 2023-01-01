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

// Send a raw request to the inverter
func (protocol *Protocol) Send(request Request) (*Response, error) {
	length, err := request.ResponseLength()
	if err != nil {
		return nil, err
	}
	protocol.mutex.Lock()
	defer protocol.mutex.Unlock()
	log.Debugf("> %s", request)
	protocol.connection.Write(request.Message())
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
	response := Response{message: data}
	log.Debugf("< %s", response)
	// TODO: validate CRC here
	return &response, err
}

// Query one or more variables
func (protocol *Protocol) Query(variables []*Variable) error {
	areas := []Area{}
	for i := 0; i < len(variables); i++ {
		areas = append(areas, (variables)[i].Area())
	}
	areas = ReduceAreas(areas)
	memories := []Memory{}
	for i := 0; i < len(areas); i++ {
		area := areas[i]
		response, err := protocol.Send(NewRequestQuery(area))
		if err != nil {
			return err
		}
		data, err := response.Message().Data()
		if err != nil {
			return err
		}
		memory := NewMemory(area)
		memory.SetData(*data)
		memories = append(memories, memory)
	}
	for i := 0; i < len(variables); i++ {
		variable := variables[i]
		memory := ExtractMemory(variable.Area(), memories)
		variable.memory.SetData(memory.Data())
	}
	return nil
}

func (protocol *Protocol) QueryOne(variable *Variable) error {
	return protocol.Query([]*Variable{variable})
}

func (protocol *Protocol) WriteOne(variable Variable) error {
	request := NewRequestWrite(variable.Memory())
	// TODO: check that response belongs to request
	_, err := protocol.Send(request)
	if err != nil {
		return err
	}
	return nil
}

func (protocol *Protocol) Login(password string) error {
	// TODO: check the comm port at login to determine if we're already logged in / how to disconnect later
	err := protocol.QueryOne(&VarLoginHash)
	if err != nil {
		return err
	}

	responseHash := CalculateLoginHash(password, VarLoginHash.Memory().Data())
	VarLoginHash.memory.SetData(responseHash)
	err = protocol.WriteOne(VarLoginHash)
	if err != nil {
		return err
	}

	err = protocol.QueryOne(&VarLoginStatus)
	if err != nil {
		return err
	}
	if !bytes.Equal(VarLoginStatus.Memory().Data(), []byte("\x01\x00")) {
		return errors.New("Invalid login status")
	}
	return nil
}

func (protocol *Protocol) Logout() error {
	// TODO: check which comm port we're actually connected to before disconnnecting port 1
	// TODO: make it possible to call .SetValue(1) here?
	VarSpLinkDisconnectingComms1.memory.SetData(Data("\x01\x00"))
	return protocol.WriteOne(VarSpLinkDisconnectingComms1)
}
