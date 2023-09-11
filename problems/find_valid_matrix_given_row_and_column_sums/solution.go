func restoreMatrix(rowSum []int, colSum []int) [][]int {

    res := make([][]int, len(rowSum))

    for i := range res {
        res[i] = make([]int, len(colSum))
    }
    for i := range rowSum {
        for j := range colSum {
            min := rowSum[i]
            if rowSum[i] > colSum[j] {
                min = colSum[j]
            } 
            rowSum[i] -= min
            colSum[j] -= min
            res[i][j] = min
        }
    }
    return res
}