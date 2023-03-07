package main

import (
	"fmt"
)

var N int = 9
var rows []int = make([]int, N)
var cols []int = make([]int, N)
var boxes []int = make([]int, N)
var finish bool = false

// 判断当前填的数是否符合
func isValidSudoku(v, r, c int) bool {
	pos := 1 << (v - 1)
	if (pos & rows[r]) > 0 {
		return false
	}
	if (pos & cols[c]) > 0 {
		return false
	}
	idx := (r/3)*3 + c/3
	if (boxes[idx] & pos) > 0 {
		return false
	}
	return true
}

func initBoard(board *[][]byte) {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if (*board)[r][c] != '.' {
				val := (*board)[r][c] - '0'
				pos := 1 << (val - 1)
				rows[r] |= pos
				cols[c] |= pos
				idx := (r/3)*3 + c/3
				boxes[idx] |= pos
			}
		}
	}
}

func solveSudoku(board [][]byte) {
	initBoard(&board)
	// 从第一层开始填
	backTrack(&board, 0, 0)
}

func place(Board *[][]byte, r, c, v int) {
	pos := 1 << (v - 1)
	rows[r] |= pos
	cols[c] |= pos
	idx := (r/3)*3 + c/3
	boxes[idx] |= pos
	(*Board)[r][c] = byte(v + '0')
}

func remove(Board *[][]byte, r, c, v int) {
	pos := 1 << (v - 1)
	rows[r] ^= pos
	cols[c] ^= pos
	idx := (r/3)*3 + c/3
	boxes[idx] ^= pos
	(*Board)[r][c] = '.'
}
func placeNext(Board *[][]byte, r, c int) {
	if r == N-1 && c == N-1 {
		finish = true
	} else {
		if c == N-1 {
			backTrack(Board, r+1, 0)
		} else {
			backTrack(Board, r, c+1)
		}
	}
}
func backTrack(Board *[][]byte, r, c int) {
	if (*Board)[r][c] == '.' {
		for i := 1; i <= N; i++ {
			if isValidSudoku(i, r, c) {
				place(Board, r, c, i)
				placeNext(Board, r, c)
				if !finish {
					remove(Board, r, c, i)
				}
			}
		}
	} else {
		placeNext(Board, r, c)
	}
}

func main() {
	board := make([][]byte, 9, 9)
	for i := 0; i < 9; i++ {
		fmt.Scanln(&board[i])
	}
	fmt.Println("*****************")
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		fmt.Printf("%s\n", board[i])
	}
}
