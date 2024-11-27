package medium

import "slices"

func longestCommonPrefix(strs []string) string {
    slices.SortFunc(strs, func(a, b string) int {
		if len(a) > len(b) {
			return 1
		} else if len(a) == len(b) {
			return 0
		}
		return -1
	})

	var common string

	for i:=0;i<len(strs[0]);i++ {
		temp := string(strs[0][i])
		for j:=1;j<len(strs);j++ {
			if temp != string(strs[j][i]) {
				goto out
			}
		}

		common += string(strs[0][i])
	}

	out:
    return common
}