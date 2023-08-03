package buffers

import (
	"bytes"
	"io"
	"testing"
)

func TestBufferWriter(t *testing.T) {
	writer := NewBufferWriter()
	writer.Write([]byte("Hello, "))
	writer.Write([]byte("World!"))
	expected := "Hello, World!"
	if writer.String() != expected {
		t.Errorf("got %v, want %v", writer.String(), expected)
	}
}

func TestBufferReader(t *testing.T) {
	bw := NewBufferWriter()

	expected := "Hello, Dave!"
	n, err := bw.Write([]byte(expected))
	if err != nil {
		t.Fatalf("failed to write to BufferWriter: %v", err)
	}

	br := NewBufferReader(bytes.NewBuffer(bw.buffer.Bytes()))

	buf := make([]byte, n)
	_, err = br.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatalf("failed to read from BufferReader: %v", err)
	}

	if got := string(buf); got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}
