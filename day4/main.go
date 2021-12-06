package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(data string) ([]int, []Board) {
	lines := strings.Split(strings.TrimSpace(data), "\n\n")

	draws := []int{}
	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(string(s))
		draws = append(draws, n)
	}

	boards := []Board{}
	for _, rawBoard := range lines[1:] {
		boards = append(boards, NewBoard(rawBoard))
	}

	return draws, boards
}

func solvePart1(data string) int {
	draws, boards := parse(data)

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

func solvePart2(data string) int {
	draws, boards := parse(data)

	winners := map[int]Board{}

	for _, draw := range draws {
		for i, _ := range boards {
			if boards[i].Has(draw) {
				boards[i].Mark(draw)
			}

			if boards[i].HasCompleteRow() || boards[i].HasCompleteColumn() {
				winners[i] = boards[i]
			}

			if len(winners) == len(boards) {
				return boards[i].ComputeScore(draw)
			}
		}
	}

	return 0
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
	fmt.Printf("solution for Part 2: %v\n", solvePart2(input))
}
