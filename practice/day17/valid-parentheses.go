package day17

func IsValid(s string) bool {
	if isClosingChar(rune(s[0])) {
		return false
	}

	var stack []rune

	for _, char := range s {
		if len(stack) > 0 {
			lastChar := stack[len(stack)-1]
			switch char {
			case '}':
				if lastChar != '{' {
					return false
				}
				stack = stack[:len(stack)-1]
			case ']':
				if lastChar != '[' {
					return false
				}
				stack = stack[:len(stack)-1]
			case ')':
				if lastChar != '(' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, char)
	}
	return true
}

func isClosingChar(char rune) bool {
	return char == '}' || char == ']' || char == ')'
}
