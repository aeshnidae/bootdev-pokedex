package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "HELLO_WORLD",
			expected: []string{"hello_world"},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test %d failed - Length does not match", i+1)
		}
		for j := range actual {
			word := actual[j]
			expectedWord := c.expected[j]
			if word != expectedWord {
				t.Errorf("Test %d failed - %v != %v", i+1, word, expectedWord)
			}
		}
	}

}
