package sp

import (
	"log"
	"net"
)

type Connection struct {
	conn *net.TCPConn
}

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

func (c *Connection) write(data []byte) {
	log.Printf("write to sp = %s", string(data))
	_, err := c.conn.Write(data)
	if err != nil {
		log.Fatalf("Write to sp failed: %s", err.Error())
	}
}

func (c *Connection) read(length int) []byte {
	reply := make([]byte, length)
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
