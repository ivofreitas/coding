package stack

import "strings"

var (
	stack, res []string
)

func GenerateParenthesis(n int) []string {
	stack, res = make([]string, 0), make([]string, 0)

	backtrack(0, 0, n)

	return res
}

func backtrack(openN, closeN, n int) {
	if n == openN && openN == closeN {
		res = append(res, strings.Join(stack, ""))
		return
	}

	if openN < n {
		stack = append(stack, "(")
		backtrack(openN+1, closeN, n)
		stack = stack[:len(stack)-1]
	}

	if closeN < openN {
		stack = append(stack, ")")
		backtrack(openN, closeN+1, n)
		stack = stack[:len(stack)-1]
	}
}
