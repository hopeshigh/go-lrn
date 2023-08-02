package person

import (
	"reflect"
	"testing"
)

func TestNewPersonInputs(t *testing.T) {
	tests := []struct {
		desc     string
		params   PersonParams
		expected *Person
		isError  bool
	}{
		{
			desc:     "empty params",
			params:   PersonParams{},
			expected: &Person{},
		},
		{
			desc:     "only first name",
			params:   PersonParams{FirstName: "Dave"},
			expected: &Person{firstName: "Dave"},
		},
		{
			desc:     "first and last name",
			params:   PersonParams{FirstName: "Dave", LastName: "NotDave"},
			expected: &Person{firstName: "Dave", lastName: "NotDave"},
		},
		{
			desc:     "first name and age",
			params:   PersonParams{FirstName: "Dave", Age: 50},
			expected: &Person{firstName: "Dave", age: 50},
		},
		{
			desc:     "all fields",
			params:   PersonParams{FirstName: "Dave", LastName: "NotDave", Age: 50},
			expected: &Person{firstName: "Dave", lastName: "NotDave", age: 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := NewPerson(tt.params)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, wanted: %v", got, tt.expected)
			}
		})
	}
}

func TestPersonGettersAndSetters(t *testing.T) {
	tests := []struct {
		desc     string
		input    *Person
		expected *Person
	}{
		{
			desc: "Empty inputs",
			input: &Person{
				firstName: "",
				lastName:  "",
				age:       0,
			},
			expected: &Person{
				firstName: "",
				lastName:  "",
				age:       0,
			},
		},
		{
			desc: "Normal inputs",
			input: &Person{
				firstName: "Dave",
				lastName:  "NotDave",
				age:       50,
			},
			expected: &Person{
				firstName: "Dave",
				lastName:  "NotDave",
				age:       50,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			p := &Person{}
			p.SetFirstName(tt.input.firstName)
			p.SetLastName(tt.input.lastName)
			p.SetAge(tt.input.age)

			if p.firstName != tt.expected.GetFirstName() || p.lastName != tt.expected.GetLastName() || p.age != tt.expected.GetAge() {
				t.Errorf("input: %v, got: %v, wanted: %v", tt.input, p, tt.expected)
			}
		})
	}
}
