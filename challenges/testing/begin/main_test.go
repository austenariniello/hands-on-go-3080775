// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"#00", 0},
		{"a2 32 ^ &/", 1},
		{"812 %6ab//", 2},
	}

	lc := letterCounter{}

	// range over tests and treat as subtests
	for _, tc := range tests{
		t.Run(tc.input, func(t *testing.T){
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}

}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"#00", 2},
		{"a2 32 ^ &/", 3},
		{"812 %6ab//", 4},
	}

	nc := numberCounter{}

	// range over tests and treat as subtests
	for _, tc := range tests{
		t.Run(tc.input, func(t *testing.T){
			if got := nc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}

}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"#00", 1},
		{"a2 32 ^ &/", 6},
		{"812 %6ab//", 4},
	}

	sc := symbolCounter{}

	// range over tests and treat as subtests
	for _, tc := range tests{
		t.Run(tc.input, func(t *testing.T){
			if got := sc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}

}
