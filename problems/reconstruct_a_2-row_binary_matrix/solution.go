func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
    matrix := [][]int{}
    matrix = append(matrix, []int{})
    matrix = append(matrix, []int{})
    for i := range colsum {
        if colsum[i] > upper + lower { return [][]int{} }
        matrix[0] = append(matrix[0], 0)
        matrix[1] = append(matrix[1], 0)
        if upper <= lower && colsum[i] == 1{
            
            lower --
            matrix[1][len(matrix[1]) - 1] = 1
        } else if lower < upper && colsum[i] == 1{
            
            upper --
            fmt.Println(upper < 0, "1", i)
            matrix[0][len(matrix[0]) - 1] = 1
        } else {
            if colsum[i] == 2 {
                if upper == 0 && lower == 0 {
                    return [][]int{}
                }
                upper --
                lower --
                fmt.Println(upper< 0, "3",i)
                matrix[0][len(matrix[0]) - 1] = 1
                matrix[1][len(matrix[1]) - 1] = 1
            } 
        }
    }
    if upper >= 1 || lower >= 1 {
        return [][]int{}
    }
    return matrix
}

/*
5, 5, 

2,1,2,0,1,0,1,2,0,1
4 3 2 2 1 1 0 
4 4 3 3 3 3 3

[1,1,1,0,1,0,0,1,0,0]
[1,0,1,0,0,0,1,1,0,1]
*/





