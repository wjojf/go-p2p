package core

// Transport represents a network transport between the nodes in the network.
// This can be of the form (TCP, UDP, ws, etc...)
type Transport interface {
	ListenAndAccept() error
	Hub() <-chan Message
	Addr() string
	Dial(string) error
	Close() error
}
