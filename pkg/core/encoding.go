package core

import (
	"errors"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type DefaultDecoder struct{}

func (d DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("empty message")
	}

	msg.Payload = buf[:n]

	return nil
}
