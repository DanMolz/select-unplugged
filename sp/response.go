package sp

import "log"

// Response message from SP Pro.
type Response struct {
	request Request
	message Message
}

func NewResponse(request Request) Response {
	return Response{request: request}
}

func (r Response) ExpectedLength() int {
	requestType := string(r.request[0])
	requestLength := len(r.request)
	dataLength := (int(r.request[1]) + 1) * 2
	crcLength := 2
	if requestType == "W" {
		return requestLength
	}
	if requestType != "Q" {
		log.Fatalf("Unexpected request type %s", requestType)
	}
	return requestLength + dataLength + crcLength
}
