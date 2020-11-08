package week06

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, clo := len(grid), len(grid[0])
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, clo)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < clo; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < clo; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[row-1][clo-1]
}

func minPathSum2(grid [][]int) int {
	l := len(grid)
	if l < 1 {
		return 0
	}
	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 && i != 0 {
				grid[i][j] += grid[i-1][j]
			} else if i != 0 && j != 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[l-1][len(grid[l-1])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
