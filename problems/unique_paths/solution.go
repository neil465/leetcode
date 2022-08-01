var dp = map[[2]int]int{}

func uniquePaths(m int, n int) int {
    dp = map[[2]int]int{}
    return pathFinder(m,n)
}
func pathFinder(m int, n int) int {
    if dp[[2]int{m,n}] > 0 {
        return dp[[2]int{m,n}]
    }
    if m == 1 && n == 1{return 1}
    if m == 1{return 1}
    if n == 1{return 1}
    dp[[2]int{m,n}] = pathFinder(m-1,n)+pathFinder(m,n-1)
    return dp[[2]int{m,n}]
}