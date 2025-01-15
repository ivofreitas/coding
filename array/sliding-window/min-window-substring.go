package sliding_window

/*
*
Given two strings s and t of lengths m and n respectively, return the minimum window substring

	of s such that every character in t (including duplicates) is included in the window.

If there is no such substring, return the empty string "".
*/
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	start, end, count := 0, 0, len(t)
	charT := make([]int, 128)
	minLength := int(^uint(0) >> 1)
	startLength := 0

	for _, char := range t {
		charT[char]++
	}

	for end < len(s) {
		if charT[s[end]] > 0 {
			count--
		}
		charT[s[end]]--
		end++

		for count == 0 {
			if end-start < minLength {
				startLength = start
				minLength = end - start
			}

			if charT[s[start]] == 0 {
				count++
			}

			charT[s[start]]++
			start++
		}
	}

	if minLength == int(^uint(0)>>1) {
		return ""
	}

	return s[startLength : minLength+startLength]
}
