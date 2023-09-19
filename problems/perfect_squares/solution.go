var dp = map[int]int{}
func numSquares(n int) int {
    dp = map[int]int{}
    return re(n, 0)
}
func re(n, d int) int {

    if n == 0 {

        return d
    }
    if dp[n] > 0 {
        return dp[n] + d
    }
    k := 1000000
    for i := 1 ; i < 1000 && i * i <= n; i ++ {
        t := re(n - (i * i) , 1)
        if t < k {
            k = t
        }
    }
    dp[n] = k
    return d + k
}