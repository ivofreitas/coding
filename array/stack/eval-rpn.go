package stack

import "strconv"

func EvalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, value := range tokens {
		if intValue, ok := isNumber(value); ok {
			stack = append(stack, intValue)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, calculate(a, b, value))
		}
	}

	return stack[len(stack)-1]
}

func isNumber(a string) (int, bool) {
	value, err := strconv.Atoi(a)
	return value, err == nil
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return 0
}
