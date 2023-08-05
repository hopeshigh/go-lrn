package basic

import (
	"os"
	"testing"
)

func TestGetTodos(t *testing.T) {
	out, err := GetTodos()
	if err != nil {
		t.Fatalf("GetTodos failed with error: %v", err)
	}

	file, err := os.ReadFile("test_data.json")
	if err != nil {
		t.Fatalf("failed to read test data: %v", err)
	}

	if out != string(file) {
		t.Errorf("got %v want %v", out, string(file))
	}
}
