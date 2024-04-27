package storage

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

type PathTransformFunc func(string) string

func CASPathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	blockSize := 5

	n := len(hashString) / blockSize
	blocks := make([]string, n)

	for i := 0; i < n; i++ {
		start := i * blockSize
		blocks[i] = hashString[start : start+blockSize]
	}

	return strings.Join(blocks, "/")
}
