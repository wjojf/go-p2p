package core

import "net"

// Peer represents a remote node in the network.
type Peer interface {
	net.Conn
	Send([]byte) error
	OnStream() error
	CloseStream()
}
