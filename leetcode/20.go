package leetcode

func isValid(s string) bool {
	stack := make([]string, 0, len(s))
	brackets := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for _, val := range s {
		strVal := string(val)
		if _, ok := brackets[strVal]; ok {
			stack = append(stack, strVal)
		} else {
			lastIndex := len(stack) - 1
			if len(stack) == 0 || brackets[stack[lastIndex]] != strVal {
				return false
			}
			stack = stack[:lastIndex]
		}
	}

	return len(stack) == 0
}
