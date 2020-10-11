package week02

func nthUglyNumber(n int) int {
	return nthUglyNumberN(n, 2, 3, 5)
}

func nthUglyNumberN(n int, a int, b int, c int) int {
	nth := 1
	uglyList := make([]int, n)
	uglyList[0] = 1

	idxa := 0
	idxb := 0
	idxc := 0

	for {
		if nth == n {
			return uglyList[n-1]
		}

		xa := uglyList[idxa] * a
		xb := uglyList[idxb] * b
		xc := uglyList[idxc] * c

		uglyList[nth] = min3(xa, xb, xc)

		if uglyList[nth] == xa {
			idxa++
		}
		if uglyList[nth] == xb {
			idxb++
		}
		if uglyList[nth] == xc {
			idxc++
		}

		nth++
	}
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
