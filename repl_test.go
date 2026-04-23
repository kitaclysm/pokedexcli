package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	// declare named test cases
	cases := map[string]struct {
		input string
		expected []string
	}{
		"simple": {
			input:		"hello world funky boy",
			expected:	[]string{"hello", "world", "funky", "boy"},
		},
		"empty": {
			input:		"",
			expected:	[]string{},
		},
		"spaced": {
			input:		"   hello   world   ",
			expected:	[]string{"hello", "world"},
		},
	}

	// loop through test cases
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			want := tc.expected
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("expected: %v received: %v", want, got)
			}
		})
	}
}