package sp

import (
	"log"
	"net"
	"sync"
)

type Connection struct {
	conn  *net.TCPConn
	mutex sync.Mutex
}

var spMutex sync.Mutex

func (c *Connection) Start() {
	servAddr := "127.0.0.1:6666"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed: %s", err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Dial failed: %s", err.Error())
	}
	c.conn = conn
}

func (c *Connection) Send(data []byte) []byte {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.write(data)
	return c.read()
}

func (c *Connection) write(data []byte) {
	log.Printf("write to sp = %s", string(data))
	_, err := c.conn.Write(data)
	if err != nil {
		log.Fatalf("Write to sp failed: %s", err.Error())
	}
}

func (c *Connection) read() []byte {
	reply := make([]byte, 1024)
	_, err := c.conn.Read(reply)
	if err != nil {
		log.Fatalf("read from sp failed: %s", err.Error())
	}

	log.Printf("read from sp = %s\n", string(reply))
	return reply
}

func (c *Connection) Close() {
	c.conn.Close()
}
