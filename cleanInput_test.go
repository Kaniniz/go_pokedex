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
			input:    " Hello, World! ",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "Nothing",
			expected: []string{"nothing"},
		},
		{
			input:    "",
			expected: []string{""},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word = %s, wanted: %s", word, expectedWord)
				return
			}
		}
	}
}
