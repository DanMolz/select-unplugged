package sp

// Connection is the plain serial connection to a SP Pro.
type Connection interface {
	Open() error
	Read(buf *[]byte) (int, error)
	Write(data []byte) (int, error)
	Close() error
}
