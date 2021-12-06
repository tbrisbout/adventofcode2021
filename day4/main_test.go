package main

import "testing"

const testInput = `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`

const testRawBoard = `14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestHasCompleteRow(t *testing.T) {
	t.Run("complete row should be false", func(t *testing.T) {
		board := NewBoard(testRawBoard)

		board.Mark(10)
		board.Mark(16)
		board.Mark(9)
		board.Mark(19)
		board.Mark(12)
		board.Mark(3)
		board.Mark(20)

		if board.HasCompleteRow() {
			t.Fatal("expected board not to have complete row")
		}
	})

	t.Run("complete row should be true", func(t *testing.T) {
		board := NewBoard(testRawBoard)

		board.Mark(10)
		board.Mark(16)
		board.Mark(15)
		board.Mark(9)
		board.Mark(19)

		if !board.HasCompleteRow() {
			t.Fatal("expected board to have complete row")
		}
	})
}

func TestHasCompleteColumn(t *testing.T) {
	t.Run("complete column should be false", func(t *testing.T) {
		board := NewBoard(testRawBoard)

		board.Mark(10)
		board.Mark(16)
		board.Mark(9)
		board.Mark(19)
		board.Mark(12)
		board.Mark(3)
		board.Mark(20)

		if board.HasCompleteColumn() {
			t.Fatal("expected board not to have complete column")
		}
	})

	t.Run("complete column should be true", func(t *testing.T) {
		board := NewBoard(testRawBoard)

		board.Mark(10)
		board.Mark(24)
		board.Mark(9)
		board.Mark(19)
		board.Mark(26)
		board.Mark(3)
		board.Mark(6)

		if !board.HasCompleteColumn() {
			t.Fatal("expected board to have complete column")
		}
	})

}

func TestSolvePart1(t *testing.T) {
	want := 4512
	got := solvePart1(testInput)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// func TestSolvePart2(t *testing.T) {
// 	want := 1924
// 	got := solvePart2(testInput)

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }
