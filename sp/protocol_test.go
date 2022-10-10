package sp

import (
	"fmt"
	"sync"
	"testing"

	"github.com/neerolyte/select-unplugged/sp/mocks"
	"github.com/stretchr/testify/assert"
)

var strCalls []string
var lastWrite []byte

type ConnectionMock struct {
	Connection
}

func (c *ConnectionMock) Read(length int) []byte {
	strCalls = append(strCalls, fmt.Sprintf("Read(%d)", length))
	return lastWrite
}

func (c *ConnectionMock) Write(data []byte) {
	strCalls = append(strCalls, fmt.Sprintf("Write(%v)", data))
	lastWrite = data
}

func (c *ConnectionMock) Close() {
	strCalls = append(strCalls, "Close()")
}

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

func TestProtocolLocking(t *testing.T) {
	connection := &ConnectionMock{}

	req1 := NewRequestQuery(Area{
		address: 0x01,
		words:   1,
	})
	req2 := NewRequestQuery(Area{
		address: 0x02,
		words:   2,
	})

	protocol := NewProtocol(connection)
	var wg sync.WaitGroup
	var res1 Message
	var res2 Message
	var err1 error
	var err2 error
	wg.Add(2)
	go func() { res2, err2 = protocol.Send(req2); wg.Done() }()
	// this is fun, the second goroutine appears to win the race (reliably?)
	go func() { res1, err1 = protocol.Send(req1); wg.Done() }()
	wg.Wait()

	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
	assert.Equal(t, []byte(req1), []byte(res1))
	assert.Equal(t, []byte(req2), []byte(res2))
	assert.Equal(t, []string{
		"Write([81 1 1 0 0 0 181 83])",
		"Read(14)",
		"Write([81 2 2 0 0 0 180 107])",
		"Read(16)",
	}, strCalls)
}
