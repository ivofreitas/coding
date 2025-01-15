package sliding_window

func MaxProfit(prices []int) int {
	start, end, maxResult := 0, 1, 0

	for end < len(prices) {

		for prices[start] > prices[end] && start < len(prices)-2 {
			start++
			end = start + 1
		}

		maxResult = max(maxResult, prices[end]-prices[start])
		end++
	}

	return maxResult
}
