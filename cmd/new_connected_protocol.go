package cmd

import "github.com/neerolyte/select-unplugged/sp"

func NewConnectedProtocol() *sp.Protocol {
	// TODO: Choose connection based on root level opts (or whether a daemon is running?)
	connection := sp.ConnectionSerial{}
	connection.Open()
	return sp.NewProtocol(&connection)
}
