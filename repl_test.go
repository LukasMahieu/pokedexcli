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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " PIKACHU",
			expected: []string{"pikachu"},
		},
		{
			input:    "PikAchu! likes bulbasaur? ",
			expected: []string{"pikachu!", "likes", "bulbasaur?"},
		},
		{
			input:    "pikachu",
			expected: []string{"pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match between %s and %s", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s, got %s", expectedWord, word)
			}
		}
	}
}
