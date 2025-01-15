package sliding_window

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var (
		s1Chars    [26]int
		start, end int
	)

	for _, char := range s1 {
		s1Chars[char-'a']++
	}

	for end < len(s2) {
		s1Chars[s2[end]-'a']--
		if s1Chars == [26]int{} {
			return true
		}

		if end+1 >= len(s1) {
			s1Chars[s2[start]-'a']++
			start++
		}
		end++
	}

	return false
}
