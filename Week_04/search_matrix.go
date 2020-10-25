package week04

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	colum := len(matrix[0]) - 1
	for row := 0; row < len(matrix); row++ {
		if matrix[row][colum] >= target {
			for i := 0; i <= colum; i++ {
				if matrix[row][i] == target {
					return true
				}
			}
			return false
		}
	}
	return false
}

func searchMatrixBinarySearch(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target > matrix[m-1][n-1] || target < matrix[0][0] {
		return false
	}

	left, right := 0, m*n-1
	mid, r, c := 0, 0, 0
	for left <= right {
		mid = left + (right-left)>>1
		r = mid / n
		c = mid % n
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left >= m*n || matrix[left/n][left%n] != target {
		return false
	}
	return true
}
