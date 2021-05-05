func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	res := make([][]int, r)
	row := -1
	col := 0
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < len(mat); i++ {
		for _, i2 := range mat[i] {
			if row == len(res[0])-1 {
				row = 0
				if col != len(res)-1 {
					col++
				}
            }else{
                row++
            }
          
            
			res[col][row] = i2
		}
	}
	return res
}
