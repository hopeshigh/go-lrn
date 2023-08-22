package firstrecurringchar

import (
	"testing"
)

func TestFirstRecurringChar(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"abcdef", ""},
		{"abcdaef", "a"},
		{"abbcdef", "b"},
		{"aabbcc", "a"},
		{"", ""},
		{"a", ""},
		{"go gopher", "g"},
		{"123456789", ""},
		{"1123456789", "1"},
		{"  ", " "},
		{"!@#!@#!@#", "!"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, _ := FirstRecurringChar(test.input)
			if got != test.expect {
				t.Errorf("")
			}
		})
	}
}
