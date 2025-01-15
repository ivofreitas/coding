package two_pointers

func MaxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, area)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
