var dp [100000]int

func countBits(n int) []int {
    dp = [100000]int{}
    res := []int{}
    for i := 0 ; i <= n ; i++{
        res = append(res,bits(i))
    }
    return res
}
func bits(i int) int{
    if dp[i] > 0 {
        return dp[i]
    }
    if i == 0 {
        return 0
    }else if i == 1{
        return 1
    }else if i%2 == 0 {
        dp[i] = bits(i/2) 
        return dp[i]
    }
    dp[i] = 1+bits(i/2)
    return dp[i]
}