package trees

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"hash"
)

var hasher = sha1.New()

const ARRAY_SIZE = 100

type BloomFilter struct {
	bitfield [ARRAY_SIZE]struct{}
}

func (bf *BloomFilter) Exists(value string) bool {
	h := hashPosition(value)
	return bf.bitfield[h] == struct{}{}
}

func (bf *BloomFilter) Set(value string) {
	h := hashPosition(value)
	bf.bitfield[h] = struct{}{}
}

func createHash(h hash.Hash, value string) (int, error) {
	_, err := h.Write([]byte(value))
	if err != nil {
		return -1, err
	}
	bits := h.Sum(nil)
	buf := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buf)
	return int(result), nil
}

// using horner's method
func hashFunc(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}

func hashPosition(value string) int {
	hs, _ := createHash(hasher, value)
	if hs < 0 {
		hs = -hs // ensure positive index
	}
	return hs % ARRAY_SIZE
}
