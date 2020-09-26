package week_01

func plusOne(digits []int) []int {
	dlen := len(digits)
	if dlen == 0 {
		return []int{1}
	}

	for i := dlen - 1; i >= 0; i-- {
		if digits[i] != 9 { // 进一位不等于10 则+1返回
			digits[i]++
			return digits
		} else {
			digits[i] = 0 // 是9 就为0
		}

	}
	return append([]int{1}, digits...) // 全部为0进位+1
}
