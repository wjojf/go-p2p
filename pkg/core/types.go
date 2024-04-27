package core

type Message struct {
	From    string
	Payload []byte
	Stream  bool
}
