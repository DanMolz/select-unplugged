package sp

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
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
	spConnection := ConnectionSerial{}
	spConnection.Open()
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

		result, err := protocol.Send(read)
		if err != nil {
			log.Fatalf("TODO: %s", err.Error())
		}

		clientConnection.Write(result)
	}
	// TODO: handle clients disconnecting... but how? clientConnection.Close()
}
