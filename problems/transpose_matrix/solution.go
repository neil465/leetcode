func transpose(matrix [][]int) [][]int {
	
	//[0.0][0,1][0,2]
	//[0,0] [1,0] [2,0]
	// 1 2
	// 3 4
	// 5 6
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		res[i] = make([]int, len(matrix))
	}
    fmt.Println(res)
    
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
            
			res[j][i] = matrix[i][j]
		}
	}
	return res
}