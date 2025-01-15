package binarysearch

func Search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)/2

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}

func SearchRecursive(nums []int, target int) int {
	return binarySearchRecursive(nums, 0, len(nums)-1, target)
}

func binarySearchRecursive(nums []int, low, high, target int) int {

	if high >= low {
		middle := low + (high-low)/2

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			return binarySearchRecursive(nums, middle+1, high, target)
		}

		return binarySearchRecursive(nums, low, middle-1, target)
	}

	return -1
}
