package medium

func isValid(s string) bool {
    match := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, v := range s {
		if match[v] == 0 {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != match[v] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

    return len(stack) == 0
}