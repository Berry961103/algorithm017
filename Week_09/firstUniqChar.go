package week09

func firstUniqChar(s string) int {
	var arr [26]int
	for i, k := range s {
		arr[k-'a'] = i
	}
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		}
		arr[k-'a'] = -1
	}
	return -1
}
