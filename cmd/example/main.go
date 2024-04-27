package main

import (
	"go-p2p/pkg/core"
	"go-p2p/pkg/tcp"
	"log"
)

func main() {
	tr := tcp.NewTransport(
		tcp.TransportOpts{
			ListenAddress: ":3000",
			Handshake:     core.NoHandshake,
			Decoder:       core.DefaultDecoder{},
		},
	)

	go func() {
		for {
			msg := <-tr.Hub()
			log.Printf("received message: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
