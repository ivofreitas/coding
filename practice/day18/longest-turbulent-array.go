package day18

func MaxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	left, right, maxSequence := 0, 1, 1
	var sign string
	for right < len(arr) {
		currentSign := getSign(arr[right-1], arr[right])
		if currentSign == "=" {
			left = right
		} else if currentSign == sign {
			left = right - 1
		}
		sign = currentSign
		maxSequence = max(maxSequence, right-left+1)
		right++
	}
	return maxSequence
}

func getSign(a, b int) string {
	if a-b > 0 {
		return ">"
	} else if a-b == 0 {
		return "="
	} else {
		return "<"
	}
}
