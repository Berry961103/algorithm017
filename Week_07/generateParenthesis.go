package week07

func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}

	var generate func(int, int, string)
	generate = func(left, right int, str string) {
		if right == 0 {
			res = append(res, str)
			return
		}

		if left > 0 {
			generate(left-1, right, str+"(")
		}

		if right > left {
			generate(left, right-1, str+")")
		}
	}
	generate(n, n, "")
	return res
}
