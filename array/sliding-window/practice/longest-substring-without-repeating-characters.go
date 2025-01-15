package practice

func LengthOfLongestSubstring(s string) int {
	start, end, maxLength, charPresence := 0, 0, 0, make(map[byte]int)

	for end < len(s) {
		charPresence[s[end]]++

		for charPresence[s[end]] > 1 {
			start++
			charPresence[s[start]]--
		}

		maxLength = max(maxLength, end-start+1)
		end++
	}

	return maxLength
}
