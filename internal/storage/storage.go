package storage

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

type Opts struct {
	PathTransformFunc PathTransformFunc
}

type Storage struct {
	pathTransformFunc PathTransformFunc
}

func NewStorage(opts Opts) *Storage {
	return &Storage{
		pathTransformFunc: opts.PathTransformFunc,
	}
}

func (s *Storage) WriteStream(key string, r io.Reader) error {

	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, r)
	if err != nil {
		return err
	}

	fp, err := s.getFilePath(key, *buf)
	if err != nil {
		return err
	}

	file, err := os.Create(fp)
	if err != nil {
		return err
	}

	n, err := io.Copy(file, buf)
	if err != nil {
		return err
	}

	log.Printf("Wrote %d bytes to %s\n", n, fp)

	return nil
}

func (s *Storage) getFilePath(key string, buf bytes.Buffer) (string, error) {
	path := s.pathTransformFunc(key)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}

	fileName := md5.Sum(buf.Bytes())
	return path + "/" + hex.EncodeToString(fileName[:]), nil
}
