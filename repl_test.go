package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Test cases
	tests := []struct {
		name string
		input string
		expected []string
	}{	
		{
			name: "empty input",
			input: "",
			expected: []string{},
		},
		{
			name: "single word",
			input: "lady",
			expected: []string{"lady"},
		},
		{
			name: "two words",
			input: "the batman",
			expected: []string{"the", "batman"},
		},
		{
			name: "whitespace padding",
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			name: "whitespace left",
			input: "     i am legend",
			expected: []string{"i", "am", "legend"},
		},
		{
			name: "whitespace right",
			input: "and i am iron man     ",
			expected: []string{"and", "i", "am", "iron", "man"},
		},
		{
			name: "whitespace abnormal",
			input: "to    ee   efl        tay    ",
			expected: []string{"to", "ee", "efl", "tay"},
		},
		{
			name: "capital all",
			input: "HI MOM",
			expected: []string{"hi", "mom"},
		},
		{
			name: "capital mixed",
			input: "THE mAN IS hERE",
			expected: []string{"the", "man", "is", "here"},
		},
		{
			name: "tabs and newlines",
			input: "yeah\tbaby\nbulbasaur\t woo",
			expected: []string{"yeah", "baby", "bulbasaur", "woo"},
		},
	}

	// Loop over the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := cleanInput(tt.input)
			if len(actual) != len(tt.expected) {
				t.Errorf("cleanInput(%q) should output: %v\nInstead it output: %v", tt.input, tt.expected, actual)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("cleanInput(%q) should output: %v\nInstead it output: %v", tt.input, tt.expected, actual)
			}
		})
	}
}