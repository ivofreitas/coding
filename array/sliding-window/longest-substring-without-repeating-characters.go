package sliding_window

/*
*
Given a string s, find the length of the longest substring without repeating characters.

Notes:
  - max -> calculate outside of inner loop
  - min -> calculate inside of inner loop
*/
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	maxLength := 0
	charCount := make([]int, 128)

	for end < len(s) {
		charCount[s[end]]++

		for charCount[s[end]] > 1 {
			charCount[s[start]]--
			start++
		}

		// calculate here
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}

		end++
	}

	return maxLength
}
