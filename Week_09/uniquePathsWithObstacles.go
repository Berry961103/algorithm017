package week09

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	cl := len(obstacleGrid[0])
	dp := make([]int, cl)
	dp[0] = 1
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < cl; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[cl-1]
}
