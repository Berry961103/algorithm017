package week02

func isAnggram(s, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen != tLen {
		return false
	}

	if sLen == 0 && tLen == 0 {
		return true
	}

	stmp := make(map[rune]int)

	for _, v := range s {
		stmp[v]++

	}

	for _, b := range t {
		if _, ok := stmp[b]; !ok {
			return false
		} else {
			stmp[b]--
		}

		if stmp[b] == 0 {
			delete(stmp, b)
		}

	}
	return len(stmp) == 0
}
