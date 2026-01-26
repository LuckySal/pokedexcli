package main

import (
	"reflect"
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
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    " long    series of  words    separated  by a  variable    number  of spaces   ",
			expected: []string{"long", "series", "of", "words", "separated", "by", "a", "variable", "number", "of", "spaces"},
		},
		{
			input:    "mIxEd   CasE WORDS",
			expected: []string{"mixed", "case", "words"},
		},
	}
	for i, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("test %d: expected: %v, got: %v", i+1, c.expected, actual)
		}
	}
}
