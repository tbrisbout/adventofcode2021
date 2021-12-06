package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	numbers [][]int
	marked  []int
}

func NewBoard(rawBoard string) Board {
	b := Board{}
	for _, line := range strings.Split(rawBoard, "\n") {
		row := []int{}
		for _, s := range strings.Split(strings.TrimSpace(strings.ReplaceAll(line, "  ", " ")), " ") {
			n, _ := strconv.Atoi(string(s))
			row = append(row, n)
		}
		b.numbers = append(b.numbers, row)
	}
	return b
}

// String prints a representation of the board matching inital input
func (b Board) String() string {
	s := "\n"
	s += "┌────────────────┐\n"

	for _, row := range b.numbers {
		s += "│ "

		for _, n := range row {
			if includes(b.marked, n) {
				s += bold(fmt.Sprintf("%2d ", n))
			} else {
				s += fmt.Sprintf("%2d ", n)
			}
		}

		s += "│\n"
	}

	s += "└────────────────┘\n"
	s += fmt.Sprintf("marked: %v\n", b.marked)

	return s
}

func (b Board) Has(draw int) bool {
	for _, row := range b.numbers {
		if includes(row, draw) {
			return true
		}
	}
	return false
}

func (b *Board) Mark(draw int) {
	b.marked = append(b.marked, draw)
}

func (b Board) HasCompleteRow() bool {
	for _, row := range b.numbers {
		if includesAll(b.marked, row) {
			return true
		}
	}
	return false
}

func (b Board) HasCompleteColumn() bool {
	for i := 0; i < len(b.numbers[0]); i++ {
		column := []int{}

		for j := 0; j < len(b.numbers); j++ {
			column = append(column, b.numbers[j][i])
		}

		if includesAll(b.marked, column) {
			return true
		}
	}

	return false
}

func (b Board) ComputeScore(draw int) int {
	unmarkedSum := 0
	for _, row := range b.numbers {
		for _, n := range row {
			if !includes(b.marked, n) {
				unmarkedSum += n
			}
		}
	}
	return unmarkedSum * draw
}

// -------
// HELPERS
// -------

// includes verifies if slice a includes given item
func includes(numbers []int, x int) bool {
	for _, n := range numbers {
		if n == x {
			return true
		}
	}
	return false
}

// includesAll verifies if slice a includes all items from slice b
func includesAll(a, b []int) bool {
	for _, n := range b {
		if !includes(a, n) {
			return false
		}
	}
	return true
}

// bold wraps a string with ANSI bold characters
func bold(s string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}
