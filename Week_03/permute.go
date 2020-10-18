package week03

func permute(nums []int) (ans [][]int) {
	nLen := len(nums)
	if nLen == 0 {
		return ans
	}
	var dfs func(int)
	dfs = func(n int) {
		if n == nLen-1 {
			cur := make([]int, nLen)
			copy(cur, nums)
			ans = append(ans, cur)
			return
		}

		for i := 0; i < nLen; i++ {
			// 交换n位置和i处的值
			nums[i], nums[n] = nums[n], nums[i]
			dfs(n + 1)
			//恢复交换前状态
			nums[i], nums[n] = nums[n], nums[i]
		}
	}
	dfs(0)
	return ans
}
