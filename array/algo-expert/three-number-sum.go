package algo_expert

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	result := make([][]int, 0)

	sort.Ints(array)

	for i := 0; i < len(array); i++ {
		left := i + 1
		right := len(array) - 1
		for left < right {
			current, leftNumber, rightNumber := array[i], array[left], array[right]
			currentSum := current + leftNumber + rightNumber

			if currentSum > target {
				right--
			} else if currentSum < target {
				left++
			} else {
				result = append(result, []int{current, leftNumber, rightNumber})

				right--
				left++
			}
		}
	}

	// Write your code here.
	return result
}
