package day16

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rowsMap := [9][9]bool{}
	colsMap := [9][9]bool{}
	gridMap := [9][9]bool{}

	for row := range board {
		for col := range board[row] {
			val, err := strconv.Atoi(string(board[row][col]))
			if err != nil {
				continue
			}

			val--
			gridIdx := col/3 + (row/3)*3
			if rowsMap[row][val] || colsMap[col][val] || gridMap[gridIdx][val] {
				return false
			}
			rowsMap[row][val] = true
			colsMap[col][val] = true
			gridMap[gridIdx][val] = true
		}
	}

	return true
}
