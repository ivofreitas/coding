package two_pointers

/*
*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color
are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.
*/
func sortColors(nums []int) {
	left, right, current := 0, len(nums)-1, 0

	for current < right {
		switch nums[current] {
		case 0:
			nums[current], nums[left] = nums[left], nums[current]
			left++
		case 1:
			current++
		case 2:
			nums[current], nums[right] = nums[right], nums[current]
			right--
		}
	}
}
