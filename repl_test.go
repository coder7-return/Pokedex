package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected){
			t.Errorf("The length are not equal: %v vs %v", len(actual), len(cs.expected) )
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := cs.expected[i]
			if (word != expectedWord){
				t.Errorf("%v does not equal %v", word, expectedWord)
			}
		}
	}
}