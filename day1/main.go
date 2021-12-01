package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePart1(data string) int {
	count := 0
	arr := strings.Split(strings.TrimSpace(data), "\n")

	for i, _ := range arr {
		if i == 0 {
			continue
		}

		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[i-1])

		if a > b {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
}
