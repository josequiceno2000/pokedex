package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name string
		input string
		expected []string
	}{ 
		{
			name: "Basic Input",
			input: " hello world ",
			expected: []string{"hello", "world"},
	    },
		{
			name: "Multiple Spaces",
			input: " This    has many     spaces ",
			expected: []string{"this", "has", "many", "spaces"},
		},
		{
			name: "All Uppercase",
			input: " UPPER FOR LIFE ",
			expected: []string{"upper", "for", "life"},
		},
		{
			name: "Leading/Trailing Spaces",
			input: "      leading and trailing spaces              ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			name: "Mixed Case",
			input: " ThIS INPut HAs MaNY CASEs ",
			expected: []string{"this", "input", "has", "many", "cases"},
		},
		{
			name: "Only Spaces",
			input: "                       ",
			expected: []string{},
		},
		{
			name: "Tabs and Newlines",
			input: "\tYo\nWhat\nhappened\t",
			expected: []string{"yo", "what", "happened"},
		},	
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := cleanInput(tc.input)
			if !reflect.DeepEqual(actualOutput, tc.expected) {
				t.Errorf("cleanInput(%q) should equal: %v\nInstead got: %v", tc.input, tc.expected, actualOutput)
			}
		})
	}
}