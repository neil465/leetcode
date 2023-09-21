var dp = [1000][1000]int{}

func minSteps(n int) int {
    dp = [1000][1000]int{}
    return re(1,0, n)
}
func re(cur, buf , n int) int {

    if cur == n {
        return 0
    }
    if cur > n {
        return math.MaxInt32
    }
    if dp[cur][buf] > 0 {
        return dp[cur][buf]
    }
    k := math.MaxInt32
    if buf != 0 {
        k = re(cur + buf, buf, n) + 1
    }
    if cur != buf {
        t := re(cur, cur, n) + 1
        if t < k {
            k = t
        }
    }

    dp[cur][buf] = k
    return k
}