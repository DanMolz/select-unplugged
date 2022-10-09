package sp

// Response message from SP Pro.
type Response struct {
	request Request
	message Message
}

func NewResponse(request Request) Response {
	return Response{request: request}
}
