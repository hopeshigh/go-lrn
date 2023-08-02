package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		describe string
		input    int
		expected string
		hasError bool
	}{
		{"given a negative number", -5, "input must be a positive integer", true},
		{"given a 0", 0, "input must be a positive integer", true},
		{"given a 1", 1, "1", false},
		{"given a number divisible by 3", 3, "Fizz", false},
		{"given a number divisible by 5", 5, "Buzz", false},
		{"given a number divisible by 15", 15, "FizzBuzz", false},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got, err := Fizzbuzz(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("input: %d, expected an error but got none", tt.input)
				}
			} else if err != nil {
				t.Errorf("input: %d, unexpected error: %v", tt.input, err)
			} else {
				if got != tt.expected {
					t.Errorf("input: %d, got: %q, expected: %q", tt.input, got, tt.expected)
				}
			}
		})
	}
}
