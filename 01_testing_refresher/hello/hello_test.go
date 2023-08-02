package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got: %q, wanted: %q", got, want)
	}
}

func TestHelloYou(t *testing.T) {
	got := Hello("Dave")
	want := "Hello, Dave!"

	if got != want {
		t.Errorf("got: %q, wanted: %q", got, want)
	}
}

// Refactoring into a table
func TestHelloTable(t *testing.T) {
	// Define the test struct
	tests := []struct {
		describe string
		input    string
		expected string
	}{
		{"given no input", "", "Hello, World!"},
		{"given an input", "Dave", "Hello, Dave!"},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got := Hello(tt.input)
			if got != tt.expected {
				t.Errorf("input: %q, got: %q, wanted: %q", tt.input, got, tt.expected)
			}
		})
	}
}
