package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(data string) []string {
	return strings.Split(strings.TrimSpace(data), "\n")
}

func parseCommand(command string) (string, int) {
	c := strings.Split(command, " ")
	dir := c[0]
	n, _ := strconv.Atoi(c[1])
	return dir, n
}

func solvePart1(data string) int {
	var h, d int

	for _, command := range parse(data) {
		dir, n := parseCommand(command)

		switch dir {
		case "forward":
			h += n
		case "up":
			d -= n
		case "down":
			d += n
		}
	}

	return h * d
}

func solvePart2(data string) int {
	var h, d, a int

	for _, command := range parse(data) {
		dir, n := parseCommand(command)

		switch dir {
		case "forward":
			h += n
			d += n * a
		case "up":
			a -= n
		case "down":
			a += n
		}
	}

	return h * d
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
	fmt.Printf("solution for Part 2: %v\n", solvePart2(input))
}
