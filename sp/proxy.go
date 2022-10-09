package sp

import (
	"log"
	"net"
	"strings"
)

type Proxy struct{}

func (p Proxy) Start() {
	spConnection := Connection{}
	spConnection.Start()
	address := "127.0.0.1:7528"
	l, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

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

func (p Proxy) handleConnection(clientConnection net.Conn, spConnection Connection) {
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

		spConnection.Write(read)

		result := spConnection.Read()
		clientConnection.Write(result)
	}
	clientConnection.Close()
}
