// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "#00", 0},
		{"one", "a2 32 ^ &/", 1},
		{"two", "812 %6ab//", 2},
	}

	counter := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := counter.count(tc.input)
			if got != tc.want {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"one", "abc 1,?/", 1},
		{"two", "abc 1&8 ^_", 2},
	}

	counter := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := counter.count(tc.input)
			if got != tc.want {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"four", "abc 1,?/", 4},
		{"five", "abc 12&8 ^_", 5},
	}

	counter := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := counter.count(tc.input)
			if got != tc.want {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}

}

func TestNames(t *testing.T) {
	tests := []struct {
		name    string
		counter interface{ name() string }
		want    string
	}{
		{"letter", letterCounter{"lc"}, "lc"},
		{"number", numberCounter{"nc"}, "nc"},
		{"symbol", symbolCounter{"sc"}, "sc"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.counter.name()
			if got != tc.want {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}
