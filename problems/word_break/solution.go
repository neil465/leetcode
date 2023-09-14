var dp = [300][300]int{}

func wordBreak(s string, wordDict []string) bool {
     dp = [300][300]int{}
    d := map[string]bool{}
    for _, i := range wordDict {
        d[i]= true
    }
    return re(s, 0,-1 ,d)
}
func re(s string, i, j int, dict map[string]bool) bool {

    if j == len(s) - 1  {
        return  dict[s[i:j + 1]] 
    }
    if dp[i + 1][j + 1] > 0 {
        return dp[i + 1][j + 1]== 1
    }
    if dict[s[i:j + 1]] {

        if re(s, j + 1, j +1, dict) {
           dp[i + 1][j + 1] = 1
            return true
        } else {
            dp[i + 1][j + 1] = 2
        }
    }
    k := re(s, i, j + 1, dict)
    if k == true {
       dp[i + 1][j + 1] = 1
    } else {
        dp[i + 1][j + 1] = 2
    }

    return k
}