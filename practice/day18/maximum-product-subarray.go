package day18

func MaxProduct(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	currentMax, currentMin, overallMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax := currentMax

		currentMax = max(nums[i], max(currentMax*nums[i], currentMin*nums[i]))
		currentMin = min(nums[i], min(tempMax*nums[i], currentMin*nums[i]))

		overallMax = max(overallMax, currentMax)
	}

	return overallMax
}
