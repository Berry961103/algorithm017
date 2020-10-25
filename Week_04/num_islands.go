package week04

func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == '0' || grid[line][column] == '2' {
				continue
			}
			count++
			dfs(grid, line, column)
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'  // 相邻1的地方变为2
	dfs(grid, i+1, j) // 上
	dfs(grid, i-1, j) // 下
	dfs(grid, i, j+1) // 左
	dfs(grid, i, j-1) // 右
}
