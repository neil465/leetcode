func setZeroes(matrix [][]int)  {
    r := map[int]bool{}
    c := map[int]bool{}

    for i := range matrix {
        for j := range matrix[0] {
            if matrix[i][j] == 0 {
                r[i] = true
                c[j] = true
            }
        }
    }
    for i := range matrix {
        for j := range matrix[0] {
            if r[i] || c[j] {
                matrix[i][j] = 0
            } 
        }
    }
}