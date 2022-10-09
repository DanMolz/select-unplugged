package sp

// Connection is the plain serial connection to a SP Pro.
type Connection interface {
	Read(length int) []byte
	Write(data []byte)
	Close()
}
