package sp

import (
	"github.com/tarm/serial"
)

type ConnectionSerial struct {
	serial *serial.Port
}

// Highlight missing interface methods early
var _ Connection = (*ConnectionSerial)(nil)

func (c *ConnectionSerial) Open() error {
	config := &serial.Config{Name: "/dev/ttyUSB1", Baud: 57600}
	serial, err := serial.OpenPort(config)
	if err != nil {
		return err
	}
	c.serial = serial
	return nil
}

func (c *ConnectionSerial) Read(buf *[]byte) (int, error) {
	return c.serial.Read(*buf)
}

func (c *ConnectionSerial) Write(data []byte) (int, error) {
	return c.serial.Write(data)
}

func (c *ConnectionSerial) Close() error {
	return c.serial.Close()
}
