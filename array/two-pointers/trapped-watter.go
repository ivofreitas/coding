package two_pointers

func Trap(height []int) int {
	trappedWatter, left, right, maxL, maxR := 0, 0, len(height)-1, height[0], height[len(height)-1]

	for left < right {
		if maxL <= maxR {
			left++
			maxL = max(maxL, height[left])
			trappedWatter += positiveSub(maxL, height[left])
		} else {
			right--
			maxR = max(maxR, height[right])
			trappedWatter += positiveSub(maxR, height[right])
		}
	}

	return trappedWatter
}

func positiveSub(a, b int) int {
	if a-b < 0 {
		return 0
	}
	return a - b
}
