package main

func plusOne(digits []int) []int {
    lastInd := len(digits)-1
	digits[lastInd] += 1
	var addTents int = 0
	for i:=lastInd; i >= 0; i-- {
		digits[i] += addTents
		if digits[i] == 10 {
			digits[i] = 0
			addTents = 1
		} else {
			addTents = 0
		}
	}

	if addTents == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}