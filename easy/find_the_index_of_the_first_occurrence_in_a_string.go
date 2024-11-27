package main

func strStr(haystack string, needle string) int {
    index := -1
    for i, v := range haystack {
		if v == rune(needle[0]) && len(needle)+i <= len(haystack){
			index = i
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					index = -1
				}
			}
			if index != -1 {
				return index
			}
		}
	}
	return index
}