package tcp

import (
	"errors"
	"go-p2p/pkg/core"
	"log"
	"net"
)

type Transport struct {
	listenAddr string
	listener   net.Listener

	handShake core.HandshakeFunc

	decoder core.Decoder

	hub chan core.Message

	onPeer core.OnPeerConnectedFunc
}

func NewTransport(opts TransportOpts) *Transport {
	return &Transport{
		listenAddr: opts.ListenAddress,
		handShake:  opts.Handshake,
		decoder:    opts.Decoder,
		onPeer:     opts.OnPeer,
		hub:        make(chan core.Message),
	}
}

func (t *Transport) ListenAndAccept() (err error) {

	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return
	}

	log.Printf("[tcp] Listening on port: %s\n", t.listenAddr)

	go t.startAccept()
	return
}

func (t *Transport) Addr() string {
	return t.listener.Addr().String()
}

func (t *Transport) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	go t.handleConn(conn, true)
	return nil
}

func (t *Transport) Close() error {
	return t.listener.Close()
}

func (t *Transport) Hub() <-chan core.Message {
	return t.hub
}

func (t *Transport) startAccept() {
	for {
		conn, err := t.listener.Accept()

		// Break the loop if the listener is closed.
		if errors.Is(err, net.ErrClosed) {
			return
		}

		if err != nil {
			log.Println("[tcp] Error accepting connection: ", err)
		}

		go t.handleConn(conn, false)
	}
}

func (t *Transport) handleConn(conn net.Conn, outbound bool) {

	defer func() {
		log.Printf("[tcp] closing connection: %s\n", conn.RemoteAddr())
		if err := conn.Close(); err != nil {
			log.Printf("[tcp] error closing connection: %s\n", err)
		}
	}()

	// Create peer
	peer := NewPeer(conn, outbound)

	// Handshake
	if err := t.handShake(peer); err != nil {
		return
	}

	// Call onPeer callback
	if t.onPeer != nil {
		if err := t.onPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	t.readMessages(conn, peer)
}

func (t *Transport) readMessages(conn net.Conn, peer *Peer) {

	msg := core.Message{}
	for {

		if err := t.decoder.Decode(conn, &msg); err != nil {
			log.Printf("[tcp] error decoding message: %s\n", err)
			return
		}

		if msg.Stream {
			if err := peer.OnStream(); err != nil {
				log.Printf("[tcp] error on stream: %s\n", err)
				return
			}
			continue
		}

		msg.From = conn.RemoteAddr().String()

		t.hub <- msg
	}
}
