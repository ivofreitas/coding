package main

/*
	GIVEN AN ARRAY, RETURN THE FIRST CHARACTER THAT REPEATS

	EG.: [2, 1, 3, 2, 5, 1]
	RETURN 2
*/

// naive approach
// O(n^2) time O(1) space
func firstRecurringCharacter(input []int) int {

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return input[i]
			}
		}
	}

	return 0
}

// improved time wise approach
// O(n) time O(n) space
func firstRecurringCharacter2(input []int) int {
	seenChars := make(map[int]struct{})
	for i := 0; i < len(input); i++ {
		if _, ok := seenChars[input[i]]; ok {
			return input[i]
		}
		seenChars[input[i]] = struct{}{}
	}

	return 0
}
