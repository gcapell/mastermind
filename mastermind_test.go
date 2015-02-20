package main

import "testing"

func TestParseGuess(t *testing.T) {
	c := colors
	tests := []struct {
		s string
		g guess
	}{
		{"0000", 0},
		{"1234", guess(1*c*c*c + 2*c*c + 3*c + 4)},
	}
	for _, test := range tests {
		got := parseGuess(test.s)
		if got != test.g {
			t.Errorf("parseGuess(%s)->%s, want %s",
				test.s, got, test.g)
		}
	}
}

func TestJudge(t *testing.T) {
	tests := []struct {
		hidden, guess string
		s             score
	}{
		{"0000", "0000", score{4, 0}},
		{"1234", "4321", score{0, 4}},
		{"1234", "1354", score{2, 1}},
	}
	for _, test := range tests {
		hidden, g := parseGuess(test.hidden), parseGuess(test.guess)
		got := hidden.judge(g)
		if got != test.s {
			t.Errorf("%s.judge(%s) -> %s, want %s",
				hidden, g, got, test.s)
		}
	}
}
