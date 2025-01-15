package practice

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

	var (
		start, end                 int
		charCount                  [128]int
		count                      = len(t)
		startLength, minimumLength = 0, int(^uint(0) >> 1)
	)

	for _, char := range t {
		charCount[char]++
	}

	for end < len(s) {
		if charCount[s[end]] > 0 {
			count--
		}
		charCount[s[end]]--
		end++

		for count == 0 {
			if end-start < minimumLength {
				startLength = minimumLength
				minimumLength = end - start
			}

			if charCount[s[start]] == 0 {
				count++
			}

			charCount[s[start]]++
			start++
		}
	}

	if minimumLength == int(^uint(0)>>1) {
		return ""
	}

	return s[startLength : minimumLength+startLength]
}
