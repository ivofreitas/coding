package stack

func DailyTemperatures(temperatures []int) []int {
	res, stack := make([]int, len(temperatures)), make([]int, 0)

	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}
