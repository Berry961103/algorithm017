package week08

func totalNQueens(n int) int {
	size, count := (1<<n)-1, 0
	var solve func(int, int, int)
	solve = func(row, ld, rd int) {
		if row == size {
			count++
			return
		}

		pos := size & (^(row | ld | rd))
		for pos != 0 {
			p := pos & (-pos)
			pos -= p
			solve(row|p, (ld|p)<<1, (rd|p)>>1)
		}
	}
	solve(0, 0, 0)
	return count
}
