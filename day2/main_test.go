package main

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`

	want := 150
	got := solvePart1(input)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
