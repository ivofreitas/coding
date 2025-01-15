package binarysearch

func FindMin(nums []int) int {
	left, right, minimum := 0, len(nums)-1, 5000

	for left <= right {
		if nums[right] > nums[left] {
			minimum = min(minimum, nums[left])
			break
		}

		middle := (left + right) / 2
		minimum = min(minimum, nums[middle])

		if nums[middle] >= nums[left] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return minimum
}
