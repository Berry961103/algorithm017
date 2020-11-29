package week09

func longestPalindrome(s string) string {
	index, maxLen, size := 0, 0, len(s)
	res := ""

	for index < size {
		pointerL, pointerR := index, index
		for pointerL >= 0 && s[pointerL] == s[index] {
			pointerL--
		}
		for pointerR < size && s[pointerR] == s[index] {
			pointerR++
		}

		nextPoint := pointerR
		for pointerL >= 0 && pointerR < size && s[pointerR] == s[pointerL] {
			pointerL--
			pointerR++
		}

		tmpMaxLen := pointerR - pointerL - 1
		if tmpMaxLen > maxLen {
			res = s[pointerL+1 : pointerR]
			maxLen = tmpMaxLen
		}

		index = nextPoint
	}
	return res
}
