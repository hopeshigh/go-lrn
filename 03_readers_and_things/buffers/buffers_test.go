package buffers

import (
	"testing"
)

func TestBufferWriter(t *testing.T) {
	writer := &BufferWriter{}
	writer.Write([]byte("Hello, "))
	writer.Write([]byte("World!"))
	expected := "Hello, World!"
	if writer.String() != expected {
		t.Errorf("got %v, want %v", writer.String(), expected)
	}
}
