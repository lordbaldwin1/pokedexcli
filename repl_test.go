package main

import "testing"

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
			input:    " hello world and stuff yup ",
			expected: []string{"hello", "world", "and", "stuff", "yup"},
		},
		{
			input:    "HeLLo WorlD",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HELLOWORLD ",
			expected: []string{"helloworld"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// // check if lengths match
		if len(actual) != len(c.expected) {
			t.Errorf("error: length of actual: %d does not match length of expected: %d", len(actual), len(c.expected))
		}

		// check if each word matches
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Case %d: word: %s does not match expected: %s", i, word, expectedWord)
			}
		}
	}
}
