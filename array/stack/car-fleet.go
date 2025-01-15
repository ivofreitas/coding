package stack

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	positionSpeed := make([][]int, 0)
	stack := make([]float32, 0)

	for i := 0; i < len(position); i++ {
		positionSpeed = append(positionSpeed, []int{position[i], speed[i]})
	}

	sort.SliceStable(positionSpeed, func(i, j int) bool {
		return positionSpeed[i][0] < positionSpeed[j][0]
	})

	for i := len(positionSpeed) - 1; i >= 0; i-- {
		timeToTarget := float32(target-positionSpeed[i][0]) / float32(positionSpeed[i][1])

		stack = append(stack, timeToTarget)
		if len(stack) > 1 && timeToTarget <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack)
}
