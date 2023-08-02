package basic

import (
	"math/rand"
)

// Intention here is that Readers and Writers are a bit of a mystery, so what
// better way to understand than to build them from the ground up

type MyReader interface {
	Read(p []byte) (n int, err error)
}

type MyWriter interface {
	Write(p []byte) (n int, err error)
}

type RandomReader struct{}

func (RandomReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(rand.Intn(256))
	}
	return len(p), nil
}

type DiscardWriter struct{}

func (DiscardWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
