package sliding_window

func MaxSubArray(nums []int) int {
    maxSoFar, maxEndsHere := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        maxEndsHere = max(maxEndsHere + nums[i], nums[i])
        maxSoFar = max(maxSoFar, maxEndsHere)
    }
    return maxSoFar
}