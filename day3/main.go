package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parse(data string) []string {
	return strings.Split(strings.TrimSpace(data), "\n")
}

func isOneMostCommon(arr []string, columnIndex int) bool {
	n := 0
	for _, s := range arr {
		if hasBitAtIndex(s, columnIndex, 1) {
			n++
		}
	}

	return float64(n) >= math.Ceil(float64(len(arr))/2)
}

func getMostCommonBit(arr []string, columnIndex int) int {
	if isOneMostCommon(arr, columnIndex) {
		return 1
	}

	return 0
}

func solvePart1(data string) int {
	d := parse(data)

	var gamma, epsilon string

	for i, _ := range d[0] {
		if isOneMostCommon(d, i) {
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

func hasBitAtIndex(s string, i int, bit int) bool {
	return s[i:i+1] == strconv.Itoa(bit)
}

func filterByBit(arr []string, columnIndex int, bit int) (ret []string) {
	for _, s := range arr {
		if hasBitAtIndex(s, columnIndex, bit) {
			ret = append(ret, s)
		}
	}
	return
}

func solvePart2(data string) int {
	d := parse(data)

	for i, _ := range d[0] {
		mostCommon := getMostCommonBit(d, i)

		d = filterByBit(d, i, mostCommon)
	}

	oxygen, _ := strconv.ParseInt(d[0], 2, 32)

	d = parse(data)
	for i, _ := range d[0] {
		var leastCommon int
		if !isOneMostCommon(d, i) {
			leastCommon = 1
		}
		d = filterByBit(d, i, leastCommon)
		if len(d) == 1 {
			break
		}
	}
	co2, _ := strconv.ParseInt(d[0], 2, 32)

	return int(oxygen) * int(co2)
}

func main() {
	fmt.Printf("solution for Part 1: %v\n", solvePart1(input))
	fmt.Printf("solution for Part 2: %v\n", solvePart2(input))
}
