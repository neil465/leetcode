var dp = map[[3]int]int{}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    dp = map[[3]int]int{}
    return r(m, n, maxMove, startRow, startColumn) % 1000000007
}
func r(m, n, maxMove, startRow, startColumn int) int {
    if val, ok := dp[[3]int{startRow, startColumn, maxMove}] ; ok {
        return val 
    }
    if startRow >= m || startColumn >= n || startRow < 0 || startColumn < 0{
        return 1
    }
    if maxMove == 0 {
        return 0
    }
    dp[[3]int{startRow, startColumn, maxMove}] = r(m,n,maxMove - 1, startRow - 1, startColumn) % 1000000007+ r(m,n,maxMove - 1, startRow + 1, startColumn) % 1000000007+ r(m,n,maxMove - 1, startRow, startColumn - 1)% 1000000007 + r(m,n,maxMove - 1, startRow, startColumn + 1)% 1000000007
    return dp[[3]int{startRow, startColumn, maxMove}] 
}