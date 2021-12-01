package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePart1(data string) int {
	count := 0
	arr := strings.Split(strings.TrimSpace(data), "\n")

	for i := 1; i < len(arr); i += 1 {
		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[i-1])

		if a > b {
			count++
		}
	}

	return count
}

func solvePart2(data string) int {
	count := 0
	arr := strings.Split(strings.TrimSpace(data), "\n")

	for i := 3; i < len(arr); i += 1 {
		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[i-3])

		if a > b {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
	fmt.Printf("solution for Part 2: %v\n", solvePart2(input))
}
