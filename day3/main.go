package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(data string) []string {
	return strings.Split(strings.TrimSpace(data), "\n")
}

func solvePart1(data string) int {
	d := parse(data)

	var gamma, epsilon string

	for i, _ := range d[0] {
		n := 0
		for _, s := range d {
			if s[i] == '1' {
				n++
			}
		}

		if n > (len(d) / 2) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 32)
	e, _ := strconv.ParseInt(epsilon, 2, 32)
	return int(g) * int(e)
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
}
