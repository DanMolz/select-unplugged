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
		data, err := Message(response).Data()
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

	variables := []*Variable{&VarLoginStatus}
	err = protocol.Query(variables)
	if err != nil {
		return err
	}
	if !bytes.Equal(VarLoginStatus.Memory().Data(), []byte("\x01\x00")) {
		return errors.New("Invalid login status")
	}
	return nil
}
