package day16

type CharCount [26]int

func GroupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[CharCount][]string)

	for _, str := range strs {
		var cc CharCount
		for _, char := range str {
			cc[char-'a']++
		}
		anagramsMap[cc] = append(anagramsMap[cc], str)
	}

	result := make([][]string, 0)
	for _, group := range anagramsMap {
		result = append(result, group)
	}
	return result
}
