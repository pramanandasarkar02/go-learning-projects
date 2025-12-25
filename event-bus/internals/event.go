package internals

type Event struct {
	Offset uint64
	Key    string
	Value  []byte
}