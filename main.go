package main

import (
	"fmt"
	"strconv"
	"time"
)

func sudokuCheck(sudokuBoard [][]int, cx int, cy int, item int) bool {

	genericSudokuCheck := func(f func([][]int, int, int, int, int) bool, sudokuBoard [][]int, cx int, cy int, item int) bool {
		for n := 0; n < len(sudokuBoard); n++ {
			if f(sudokuBoard, n, cx, cy, item) {
				return true
			}
		}
		return false
	}

	boxSudokuCheck := func(sudokuBoard [][]int, cx int, cy int, item int) bool {
		for i := 0; i < (len(sudokuBoard) / 3); i++ {
			for j := 0; j < (len(sudokuBoard) / 3); j++ {
				if sudokuBoard[i+cx][j+cy] == item {
					return true
				}
			}
		}
		return false
	}

	row := func(sudokuBoard [][]int, n int, cx int, cy int, item int) bool {
		return sudokuBoard[cx][n] == item
	}

	col := func(sudokuBoard [][]int, n int, cx int, cy int, item int) bool {
		return sudokuBoard[n][cy] == item
	}

	return !genericSudokuCheck(row, sudokuBoard, cx, cy, item) && !genericSudokuCheck(col, sudokuBoard, cx, cy, item) && !boxSudokuCheck(sudokuBoard, cx-cx%3, cy-cy%3, item)
}

func sudokuFindEmpty(sudokuBoard [][]int) (int, int) {
	for n := 0; n < len(sudokuBoard); n++ {
		for m := 0; m < len(sudokuBoard[0]); m++ {
			if sudokuBoard[n][m] == 0 {
				return n, m
			}
		}
	}

	return -1, -1
}

func _solveSudoku(sudokuBoard [][]int) ([][]int, bool) {
	cx, cy := sudokuFindEmpty(sudokuBoard)

	if cx == -1 && cy == -1 {
		return sudokuBoard, true
	}

	for x := 1; x <= len(sudokuBoard); x++ {
		if sudokuCheck(sudokuBoard, cx, cy, x) {
			sudokuBoard[cx][cy] = x

			_, m := _solveSudoku(sudokuBoard)
			if m {
				return sudokuBoard, true
			}

			sudokuBoard[cx][cy] = 0
		}
	}

	return nil, false
}

func solveSudoku(sudokuBoard [][]int) [][]int {
	n, m := _solveSudoku(sudokuBoard)
	if m {
		return n
	}
	return nil
}

func printBoard(sudokuBoard [][]int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			fmt.Print(strconv.Itoa(sudokuBoard[x][y]) + " ")
		}
		fmt.Println()
	}

	fmt.Println()
}

func solve(sudokuBoard [][]int) {
	fmt.Println("Solving: ")
	printBoard(sudokuBoard)
	start := time.Now()
	sudokuBoard = solveSudoku(sudokuBoard)
	elapsed := time.Since(start)
	printBoard(sudokuBoard)
	fmt.Println(elapsed)
}

func main() {
	/*
	sudokuBoard := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	solve(sudokuBoard)

	sudokuBoard = [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solve(sudokuBoard)

	
			[5, 3, 4, 6, 7, 8, 9, 1, 2],
		    [6, 7, 2, 1, 9, 5, 3, 4, 8],
		    [1, 9, 8, 3, 4, 2, 5, 6, 7],
		    [8, 5, 9, 7, 6, 1, 4, 2, 3],
		    [4, 2, 6, 8, 5, 3, 7, 9, 1],
		    [7, 1, 3, 9, 2, 4, 8, 5, 6],
		    [9, 6, 1, 5, 3, 7, 2, 8, 4],
		    [2, 8, 7, 4, 1, 9, 6, 3, 5],
		    [3, 4, 5, 2, 8, 6, 1, 7, 9]
	*/

	//

	/*sudokuBoard := [][]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}

	solveSudoku(sudokuBoard)
	*/

	sudokuBoard := [][]int{
		{6, 3, 0, 0, 0, 0, 0, 8, 1},
		{0, 2, 0, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 1, 7, 4, 3, 0},
		{0, 9, 6, 4, 0, 0, 5, 7, 0},
		{0, 0, 0, 7, 6, 2, 0, 0, 0},
		{0, 8, 0, 0, 0, 0, 6, 0, 0},
		{0, 6, 0, 0, 2, 0, 0, 0, 0},
		{3, 0, 9, 0, 0, 0, 0, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}

	solve(sudokuBoard)
}
