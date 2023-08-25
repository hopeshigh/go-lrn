package maplookuppattern

import (
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		data     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{1, 2, 3, 4, 5, 2}, true},
		{[]int{}, false},
		{[]int{100, 200}, false},
		{[]int{100, 100}, true},
	}

	for _, tt := range tests {
		got := HasDuplicate(tt.data)
		if got != tt.expected {
			t.Errorf("want: %v, got: %v", tt.expected, got)
		}
	}
}

func TestIsString(t *testing.T) {
	tests := []struct {
		desc     string
		val      interface{}
		expected bool
	}{
		{"when given a string", "hello", true},
		{"when given an integer", 42, false},
		{"when given a float", 3.14, false},
		{"when given a byte slice", []byte("hello"), false},
		{"when given nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for _, tt := range tests {
				got := IsString(tt.val)
				if got != tt.expected {
					t.Errorf("want: %v, got: %v", tt.expected, got)
				}
			}
		})
	}
}
