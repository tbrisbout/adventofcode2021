package main

import "testing"

const testInput = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`

func TestSolvePart1(t *testing.T) {
	want := 150
	got := solvePart1(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	want := 900
	got := solvePart2(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
