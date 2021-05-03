func isToeplitzMatrix(matrix [][]int) bool {
	rowlen := len(matrix[0])
	numofcollums := len(matrix)
	for i := 0; i < numofcollums-1; i++ {
		for j := 0; j < rowlen ; j++ {
			if j+1 < rowlen && matrix[i][j] != matrix[i+1][j+1]  {
				return false
			}
		}
	}
	return true
}