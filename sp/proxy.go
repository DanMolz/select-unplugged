package sp

import (
	"log"
	"net"
	"strings"
)

type Proxy struct{}

func (p Proxy) Start(address string) {
	log.Printf("Listening on %s", address)
	l, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Printf("Connecting to SP Pro")
	spConnection := ConnectionTcp{}
	spConnection.Start()
	protocol := NewProtocol(&spConnection)

	for {
		clientConnection, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Print("Accepted connection")
		go p.handleConnection(clientConnection, protocol)
	}
}

func (p Proxy) handleConnection(clientConnection net.Conn, protocol *Protocol) {
	log.Printf("Serving %s\n", clientConnection.RemoteAddr().String())
	for {
		read := make([]byte, 1024)
		_, err := clientConnection.Read(read)
		if err != nil {
			log.Fatal(err)
		}

		stringRead := strings.TrimSpace(string(read))
		log.Printf(
			"Read from %s: %s",
			clientConnection.RemoteAddr().String(),
			stringRead,
		)

		result := protocol.Send(read)

		clientConnection.Write(result)
	}
	// TODO: handle clients disconnecting... but how? clientConnection.Close()
}
