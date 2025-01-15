package practice

/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number 
of days you have to wait after the ith day to get a warmer temperature. 
If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/
func dailyTemperatures(temperatures []int) []int {

	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}
	
	for i := n - 1; i >= 0; i-- {

		// check if current temperature is higher from previous analized ones
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			// pop
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}

		//push
		stack = append(stack, i)
		
	} 

	return result
	
}