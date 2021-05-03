func diagonalSum(mat [][]int) int {
	res := 0
	col :=  len(mat[0])-1
	for i := 0; i < len(mat); i++ {
		res += mat[i][i]
		if i != col {
			res += mat[i][col]
		}

		col--
	}
	return res

}