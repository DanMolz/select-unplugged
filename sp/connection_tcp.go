package sp

import (
	"net"

	"github.com/pkg/errors"
)

type ConnectionTcp struct {
	conn *net.TCPConn
}

// Highlight missing interface methods early
var _ Connection = (*ConnectionTcp)(nil)

func (c *ConnectionTcp) Open() error {
	servAddr := "127.0.0.1:6666"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		return errors.Wrap(err, "ResolveTCPAddr failed")
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return errors.Wrap(err, "Dial failed")
	}
	c.conn = conn
	return nil
}

func (c *ConnectionTcp) Read(buf *[]byte) (int, error) {
	return c.conn.Read(*buf)
}

func (c *ConnectionTcp) Write(data []byte) (int, error) {
	return c.conn.Write(data)
}

func (c *ConnectionTcp) Close() error {
	return c.conn.Close()
}
