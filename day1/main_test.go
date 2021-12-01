package main

import "testing"

const testInput = `
199
200
208
210
200
207
240
269
260
263
	`

func TestSolvePart1(t *testing.T) {
	want := 7
	got := solvePart1(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	want := 5
	got := solvePart2(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
