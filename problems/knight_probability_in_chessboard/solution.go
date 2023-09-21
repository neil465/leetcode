var dp = [26][26][101]float64{}

func knightProbability(n int, k int, row int, column int) float64 {
    dp = [26][26][101]float64{}
    return re(row, column, n, k)
}
func re(r, c, n , k int) float64 {
    
    if r < 0 || r >= n || c< 0 || c>= n {
        return 0.0
    }
    if k == 0 {
        return 1.0
    }
    if dp[r][c][k] > 0 {
        return dp[r][c][k]
    }
    s := 0.0
    for _, m := range [][]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}} {
        s += re(r + m[0], c + m[1], n, k - 1)
    }
    dp[r][c][k] = s / 8.0
    return s / 8.0
}