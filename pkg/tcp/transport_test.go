package tcp

import (
	"github.com/stretchr/testify/assert"
	core2 "go-p2p/pkg/core"
	"testing"
)

func TestTransport(t *testing.T) {

	tr := NewTransport(
		TransportOpts{
			ListenAddress: ":3000",
			Handshake:     core2.NoHandshake,
			Decoder:       core2.DefaultDecoder{},
		},
	)

	assert.Nil(t, tr.ListenAndAccept())
}
