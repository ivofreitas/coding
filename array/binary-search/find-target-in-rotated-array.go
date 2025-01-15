package binarysearch

func SearchTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (right + left) / 2

		if nums[middle] == target {
			return middle
		}

		// left sorted portion
		if nums[middle] >= nums[left] {
			if nums[middle] < target || nums[left] > target {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else { // right sorted portion
			if nums[middle] > target || nums[right] < target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}
