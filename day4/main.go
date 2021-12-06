package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(data string) []string {
	return strings.Split(strings.TrimSpace(data), "\n\n")
}

func solvePart1(data string) int {
	d := parse(data)

	draws := []int{}
	for _, s := range strings.Split(d[0], ",") {
		n, _ := strconv.Atoi(string(s))
		draws = append(draws, n)
	}

	boards := []Board{}
	for _, rawBoard := range d[1:] {
		boards = append(boards, NewBoard(rawBoard))
	}

	for _, draw := range draws {
		for i, _ := range boards {
			if boards[i].Has(draw) {
				boards[i].Mark(draw)
			}
			if boards[i].HasCompleteRow() || boards[i].HasCompleteColumn() {
				return boards[i].ComputeScore(draw)
			}
		}

	}

	return 0
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
}
