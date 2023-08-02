package stringwriter

import (
	"fmt"
	"testing"
)

func TestStringConcatWriter(t *testing.T) {
	writer := &StringConcatWriter{}
	data := []byte("Hello, world!")
	n, err := writer.Write(data)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if n != len(data) {
		t.Errorf("expected %d bytes written, got %d", len(data), n)
	}
	if writer.String() != string(data) {
		t.Errorf("expected %s, got %s", string(data), writer.String())
	}
	fmt.Println(writer.String())
}
