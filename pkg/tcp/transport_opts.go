package tcp

import (
	core2 "go-p2p/pkg/core"
)

type TransportOpts struct {
	ListenAddress string
	Handshake     core2.HandshakeFunc
	Decoder       core2.Decoder
	OnPeer        core2.OnPeerConnectedFunc
}
