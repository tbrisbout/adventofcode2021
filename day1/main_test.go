package main

import "testing"

func TestSolvePart1(t *testing.T) {
	testInput := `
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

	want := 7
	got := solvePart1(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
