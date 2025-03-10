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
			input:    " hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "! h!!!!ello World ",
			expected: []string{"!", "h!!!!ello", "world"},
		},
		{
			input:    "suoyi zanshi   jiang",
			expected: []string{"suoyi", "zanshi", "jiang"},
		},
		{
			input:    "yoyo! hello World ",
			expected: []string{"yoyo!", "hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice length not equal: got %d, want %d", len(actual), len(c.expected))
		}
		for i := range actual {
			w := actual[i]
			expectedW := c.expected[i]

			if w != expectedW {
				t.Errorf("word not equal at actual[%d]: got %s, want %s", i, w, expectedW)
			}

		}

	}

}
