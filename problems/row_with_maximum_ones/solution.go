func rowAndMaximumOnes(mat [][]int) []int {
    row, sum := 0, 0 
    for i := range mat {
        t := 0
        for j := range mat[0] {
            t += mat[i][j]
        }
        if t > sum {
            row = i
            sum = t
        }
    }
    return []int{row, sum}
}