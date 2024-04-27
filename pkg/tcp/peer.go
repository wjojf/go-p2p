package tcp

import (
	"net"
	"sync"
)

// Peer represents a remote node over a TCP established connection.
type Peer struct {
	// The underlying connection.
	net.Conn

	// if false, the connection was initiated by the remote node.
	// if true, the connection was initiated by this peer.
	outbound bool

	// wg is used to wait for the stream to close.
	wg *sync.WaitGroup
}

func NewPeer(conn net.Conn, outbound bool) *Peer {
	return &Peer{
		Conn:     conn,
		outbound: outbound,
	}
}

func (p *Peer) OnStream() error {
	p.wg.Add(1)
	p.wg.Wait()
	return nil
}

func (p *Peer) CloseStream() {
	p.wg.Done()
}

func (p *Peer) Send(b []byte) error {
	_, err := p.Conn.Write(b)
	return err
}
