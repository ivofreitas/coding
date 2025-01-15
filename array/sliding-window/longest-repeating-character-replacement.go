package sliding_window

func CharacterReplacement(s string, k int) int {
	start, end, charCount, maxWindow, mostFreq := 0, 0, [26]int{}, 0, 0

	for end < len(s) {
		charCount[s[end]-'A']++
		if (end - start + 1 - max(mostFreq, charCount[s[end]-'A'])) > k {
			charCount[s[start]-'A']--
			start++
		}
		maxWindow = max(maxWindow, end-start+1)
		end++
	}
	return maxWindow
}
