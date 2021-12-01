package main

import (
	"fmt"
	"testing"
)

const testInput = `
199
200
208
210
200
207
240
269
260
263
	`

func TestCountIncrements(t *testing.T) {
	var tests = []struct{ want, gap int }{
		{7, 1},
		{5, 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("should return %d with gap %d", test.want, test.gap), func(t *testing.T) {
			got := countIncrements(test.gap, parse(testInput))

			if got != test.want {
				t.Errorf("got %q want %q", got, test.want)
			}
		})
	}

}
