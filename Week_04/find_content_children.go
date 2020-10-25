package week04

import (
	"sort"
)

func findContentChildren(g, s []int) int {

	gLen, sLen := len(g), len(s)
	if gLen == 0 || sLen == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0

	for i < gLen && j < sLen {
		if g[i] <= s[j] {
			i++
		}
		j++
	}

	return i
}
