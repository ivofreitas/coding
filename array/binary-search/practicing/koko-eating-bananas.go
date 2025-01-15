package practicing

import "math"

func MinEatingSpeed(piles []int, h int) int {
	left, right := 1, maxValue(piles)

	for left < right {
		k := (right + left) / 2

		if kWorks(piles, h, k) {
			right = k
		} else {
			left = k + 1
		}
	}

	return right
}

func kWorks(piles []int, h, k int) bool {
	var hours int
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile) / float64(k)))
	}
	return hours <= h
}

func maxValue(piles []int) int {
	var result int
	for _, value := range piles {
		result = max(result, value)
	}
	return result
}
