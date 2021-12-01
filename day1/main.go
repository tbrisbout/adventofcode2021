package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(data string) []string {
	return strings.Split(strings.TrimSpace(data), "\n")
}

func countIncrements(gap int, arr []string) int {
	count := 0

	for i := gap; i < len(arr); i += 1 {
		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[i-gap])

		if a > b {
			count++
		}
	}

	return count
}

func main() {
	arr := parse(input)

	fmt.Printf("solution for Part 1: %v\n", countIncrements(1, arr))
	fmt.Printf("solution for Part 2: %v\n", countIncrements(3, arr))
}
