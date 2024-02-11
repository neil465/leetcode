func modifiedMatrix(matrix [][]int) [][]int {
    cols := []int{}
    for c := range matrix[0] {
        m := -1
        for r := range matrix {
            m = max(m, matrix[r][c])
        }
        cols = append(cols, m)
    }

    for r := range matrix {
        for c := range matrix[0] {
            if matrix[r][c] == -1 {
                matrix[r][c] = cols[c]
            }
        }
    }
    return matrix
}