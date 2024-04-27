package storage

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestStorage(t *testing.T) {

	testOpts := Opts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStorage(testOpts)

	data := bytes.NewReader([]byte("some data"))

	assert.Nil(t, s.WriteStream("static", data))
}

func TestPathTransformFunc(t *testing.T) {
	key := "static"
	log.Println(CASPathTransformFunc(key))
}
