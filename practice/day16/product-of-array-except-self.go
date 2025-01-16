package day16

func ProductExceptSelf(nums []int) []int {
	answers := make([]int, len(nums))
	answers[0] = 1
	for i := 1; i < len(nums); i++ {
		answers[i] = answers[i-1] * nums[i-1]
	}

	aux := nums[len(nums)-1]
	for i := len(answers) - 2; i >= 0; i-- {
		answers[i] *= aux
		aux *= nums[i]
	}

	return answers
}
