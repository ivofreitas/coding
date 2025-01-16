package day16

func TopKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int)
	groupCount := make([][]int, 0)

	for _, num := range nums {
		numCount[num]++
	}

	for num, count := range numCount {
		groupCount[count] = append(groupCount[count], num)
	}

	result := make([]int, 0)
	for i := len(groupCount) - 1; i >= 0; i-- {
		if len(groupCount[i]) > 0 {
			j := 0
			for len(result) < k && j < len(groupCount[i]) {
				result = append(result, groupCount[i][j])
				j++
			}
		}
	}

	return result
}
