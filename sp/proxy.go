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
	spConnection := &Connection{}
	spConnection.Start()

	for {
		clientConnection, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Print("Accepted connection")
		go p.handleConnection(clientConnection, spConnection)
	}
}

func (p Proxy) handleConnection(clientConnection net.Conn, spConnection *Connection) {
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
		if stringRead == "STOP" {
			break
		}

		result := spConnection.Send(read)

		clientConnection.Write(result)
	}
	clientConnection.Close()
}
