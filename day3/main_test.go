package main

import "testing"

const testInput = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func TestSolvePart1(t *testing.T) {
	want := 198
	got := solvePart1(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	want := 230
	got := solvePart2(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
