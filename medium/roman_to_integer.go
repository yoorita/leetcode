package medium

func romanToInt(s string) int {
    number := 0
	prev := 0

	represented := map[string]int {
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := len(s)-1; i > -1; i-- {
		transformed := represented[string(s[i])]

		if prev <= transformed {
			number += transformed
		} else {
			number -= transformed
		}

		prev = transformed
	}

	return number
}