package day16

import "sort"

func fourSum(nums []int, target int) [][]int {
	result, quad := make([][]int, 0), make([]int, 0)
	sort.Ints(nums)
	kSum(nums, quad, &result, 4, 0, target)
	return result
}

func kSum(nums, quad []int, result *[][]int, k, start, target int) {

	// base case k = 2 (two sum)
	if k == 2 {
		left, right := start, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				*result = append(*result, append(quad, nums[left], nums[right]))

				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}

		return
	}

	// k != 2
	for i := start; i < len(nums)-k+1; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		quad = append(quad, nums[i])
		kSum(nums, quad, result, k-1, i+1, target-nums[i])
		quad = quad[:len(quad)-1]
	}
}
