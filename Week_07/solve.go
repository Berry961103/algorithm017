package week07

func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return // 这个矩阵最小也得是3x3的，否则不可能出现被包围的O
	}

	r, c := len(board), len(board[0]) // r -> row, c -> column
	for i := 0; i < r; i++ {          // 从第一行到最后行
		dfs(board, r, c, i, 0)   // 每行首列的每个元素 -
		dfs(board, r, c, i, c-1) // 最后一列的每个元素 -
	}

	for i := 1; i < c-1; i++ { // 从第二列到倒数第二列
		dfs(board, r, c, 0, i)   // 第一行的每个元素 |
		dfs(board, r, c, r-1, i) // 最后一行的每个元素 |
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			switch board[i][j] {
			case 'A':
				board[i][j] = 'O' //
			case 'O':
				board[i][j] = 'X' // O变成X，因为不和边界相连
			}
		}
	}
}

func dfs(board [][]byte, r, c, x, y int) {
	if x < 0 || x >= r || y < 0 || y >= c || board[x][y] != 'O' { // 越界或者这个元素不为O
		return
	}
	board[x][y] = 'A' // 没越界并且这个元素是O变成A，

	dfs(board, r, c, x-1, y) // 上
	dfs(board, r, c, x+1, y) // 下
	dfs(board, r, c, x, y-1) // 左
	dfs(board, r, c, x, y+1) // 右
}
