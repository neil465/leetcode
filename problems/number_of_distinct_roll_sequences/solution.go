var dp [100000][7][7]int
func distinctSequences(n int) int {
    dp = [100000][7][7]int{}
    return find(n,0,0)
    
}

func find(n, prev, pprev int) int{
    if n == 0 {
        return  1
    }
    if dp[n][prev][pprev] == 0 {
        for i := 1 ; i < 7 ; i++{
            if  i != prev && i != pprev&& (gcd(i, prev) == 1 || prev == 0) {
                dp[n][prev][pprev] = (dp[n][prev][pprev] + find(n-1, i, prev)) % 1000000007
            }
    }
    
    }
    return dp[n][prev][pprev]
    
}
func gcd(a int,b int) int{
    gcdnum := 0
 
    for i := 1; i <= a && i <= b ; i++ {
        if a% i == 0 && b % i == 0 {
            gcdnum = i
        } 
    }
    
    return gcdnum
}