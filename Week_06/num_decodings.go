package week06

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	} else if s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-2] == '1' || s[i-2] == '2' && s[i-1] < '7' {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func numDecodings2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	} else if s[0] == '0' {
		return 0
	}

	pre, cur, tmp := 1, 1, 0
	for i := 1; i < n; i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		case s[i] == '0':
			cur = pre
		case (s[i] < '7' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			tmp = cur
			cur = cur + pre
			pre = tmp
		default:
			pre = cur
		}
	}
	return cur
}
