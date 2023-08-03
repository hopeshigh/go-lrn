package filebuffer

import (
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	data := []byte("Hello, World!")

	tmpfile, err := os.CreateTemp("", "example.txt")
	defer os.Remove(tmpfile.Name())

	if err != nil {
		t.Fatalf("failed to create tmpfile: %v", err)
	}

	err = WriteToFile(tmpfile.Name(), data)
	if err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	readData, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("failed to read from file: %v", err)
	}

	if string(readData) != string(data) {
		t.Errorf("got %v, want %v", string(readData), string(data))
	}
}

func TestReadFromFile(t *testing.T) {
	data := []byte("Hello, Frank!")

	tmpfile, err := os.CreateTemp("", "example.txt")
	defer os.Remove(tmpfile.Name())

	if err != nil {
		t.Fatalf("failed to create tmpfile: %v", err)
	}

	err = os.WriteFile(tmpfile.Name(), data, 0644)
	if err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	readData, err := ReadFromFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("failed to read from file: %v", err)
	}

	if string(readData) != string(data) {
		t.Errorf("got %v, wants %v", string(readData), string(data))
	}
}
