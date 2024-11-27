package main

import "strings"

func lengthOfLastWord(s string) int {
    trimmed := strings.TrimSpace(s)
	lastSpaceInd := strings.LastIndex(trimmed, " ")
    return len(trimmed[lastSpaceInd+1:])
}