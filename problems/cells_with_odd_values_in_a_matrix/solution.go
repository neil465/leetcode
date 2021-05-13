func oddCells(m int, n int, indices [][]int) int {
	res := 0
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[indices[i][0]][j]++
		}
		for j := 0; j < len(matrix); j++ {
			matrix[j][indices[i][1]]++
		}
	}
	for i := 0; i < len(matrix); i++ {
		for _, i3 := range matrix[i] {
			if i3%2 != 0 {
				res++
			}
		}
	}
	return res
}