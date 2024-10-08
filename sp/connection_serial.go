package sp

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type ConnectionSerial struct {
	serial   *serial.Port
	portName string
}

// Highlight missing interface methods early
var _ Connection = (*ConnectionSerial)(nil)

func NewConnectionSerial(port string) ConnectionSerial {
	return ConnectionSerial{
		portName: port,
	}
}

func (c *ConnectionSerial) Open() error {
	config := &serial.Config{
		Name:        c.portName,
		Baud:        57600,
		ReadTimeout: time.Millisecond * 100,
		Size:        8,
		Parity:      0,
		StopBits:    1,
	}

	log.Debugf("Opening serial port %s", c.portName)
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
