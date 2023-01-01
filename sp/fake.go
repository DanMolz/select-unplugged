package sp

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type Fake struct{}

func (f Fake) Start(address string) {
	log.Printf("Listening on %s", address)
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
		go f.handleConnection(clientConnection)
	}
}

func (f Fake) handleConnection(clientConnection net.Conn) {
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
			len(request.Message()),
		)

		// TODO: handle other data here
		data := Message("\x01\x00")
		response, err := NewResponse(*request, data)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(
			"Write to %s: %s (%d)",
			clientConnection.RemoteAddr().String(),
			response,
			len(response.Message()),
		)
		clientConnection.Write(response.Message())
	}
}
