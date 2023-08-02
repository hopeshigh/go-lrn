package basic

import (
	"testing"
)

func TestRandomReader(t *testing.T) {
	reader := RandomReader{}
	data := make([]byte, 10)
	n, err := reader.Read(data)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if n != 10 {
		t.Errorf("expected 10 bytes read, got %d", n)
	}
}

func TestDiscardWriter(t *testing.T) {
	writer := DiscardWriter{}
	data := make([]byte, 10)
	n, err := writer.Write(data)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if n != 10 {
		t.Errorf("expected 10 bytes written, got %d", n)
	}
}
