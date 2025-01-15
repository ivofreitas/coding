package practice

/*
*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Notes:
  - max -> calculate it outside inner loop
  - min -> calculate it inside of inner loop
*/
func minSubArrayLen(target int, nums []int) int {

	start, end, minLength, sum := 0, 0, int(^uint(0)>>1), 0

	for end < len(nums) {
		sum += nums[end]

		for sum >= target {
			if minLength > end - start + 1 {
				minLength = end - start + 1
			}

			sum -= nums[start]
			start++
		}

		end++
	}

	if minLength == int(^uint(0)>>1) {
		return 0
	}

	return minLength
}