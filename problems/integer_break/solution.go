

func integerBreak(n int) int {
    k := re(n, 0, map[int]int{} )
    return k
}
func re(left, depth int, dp map[int]int) int {
    if left == 0 {
        if depth == 1 {
            return -1
        }
        return 1
    }
    if dp[left] > 0 {
        return dp[left]
    }
    max := 0

    for i := 1; i <= left; i++ {
        r := i * re(left - i, depth + 1, dp)
        if r > max {
            max = r
        }
    }
    dp[left] = max
    return max 
}