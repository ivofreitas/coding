package practicing

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix[0]), len(matrix)
	left, right := 0, m*n-1
	for left <= right {
		middle := (right + left) / 2
		middleElement := matrix[middle/m][middle%m]

		if target > middleElement {
			left = middle + 1
		} else if target < middleElement {
			right = middle - 1
		} else {
			return true
		}
	}

	return false
}
