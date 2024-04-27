package core

// HandshakeFunc is a function that performs a handshake with a remote peer.
type HandshakeFunc func(Peer) error

// NoHandshake is a no-op handshake function.
func NoHandshake(Peer) error {
	return nil
}
