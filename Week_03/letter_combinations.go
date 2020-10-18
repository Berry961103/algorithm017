package week03

func letterCombinations(digits string) (res []string) {
	if len(digits) == 0 {
		return nil
	}
	phone := make(map[string]string, 8) /// 列出按键
	phone["2"] = "abc"
	phone["3"] = "def"
	phone["4"] = "ghi"
	phone["5"] = "jkl"
	phone["6"] = "mno"
	phone["7"] = "pqrs"
	phone["8"] = "tuv"
	phone["9"] = "wxyz"

	var dfs func(string, int)
	dfs = func(s string, i int) {
		if i == len(digits) { // 数字的长度就是组合的长度
			res = append(res, s)
			return
		}

		d := string(digits[i])              // 取出输入的字符串中第i个数字
		letters := phone[d]                 // 数字对应的字母
		for j := 0; j < len(letters); j++ { // 根据字母个数遍历递归
			slet := string(letters[j])
			dfs(s+slet, i+1) // 用递归去遍历所有能组成的可能
		}
	}
	dfs("", 0)
	return res
}
