func mySqrt(x int) int {
    ans    := float64(x)
    delta  := 0.99999
    for math.Abs(math.Pow(ans, 2) - float64(x)) > delta {
        ans = (ans + float64(x) / ans) / 2
    }
    return int(ans)
}