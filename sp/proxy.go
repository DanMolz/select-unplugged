package sp

import (
	"net"

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

		request, err := extractRequest(read)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(
			"Read from %s: %s (%d)",
			clientConnection.RemoteAddr().String(),
			request,
			len(request),
		)

		response, err := protocol.Send(request)
		if err != nil {
			log.Fatalf("TODO: %s", err.Error())
		}

		log.Printf(
			"Write to %s: %s (%d)",
			clientConnection.RemoteAddr().String(),
			response,
			len(response),
		)
		clientConnection.Write(response)
	}
	// TODO: handle clients disconnecting... but how? clientConnection.Close()
}

func extractRequest(read []byte) (Request, error) {
	length, err := CalculateRequestLength(read)
	if err != nil {
		return nil, err
	}

	return Request(read[0:*length]), nil
}
