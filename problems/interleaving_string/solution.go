var dp = [200][200]int{}

func isInterleave(s1 string, s2 string, s3 string) bool {
    dp = [200][200]int{}
    return re(s1, s2, s3)
}
func re( s1, s2, s3 string) bool {

    if len(s3) == 0 {
        return len(s1) == 0 && len(s2) == 0
    }
    if dp[len(s1)][len(s2)] > 0 {
        return dp[len(s1)][len(s2)] == 1
    }
    
    res := false
    if len(s1) > 0 && s1[0] == s3[0]{
        res = re(s1[1:], s2, s3[1:]) 
    } 
    if len(s2) > 0 && s2[0] == s3[0] {
        res = re(s1, s2[1:], s3[1:]) || res
    }
    if res {
        dp[len(s1)][len(s2)] = 1
    } else {
        dp[len(s1)][len(s2)] = 2
    }
    return res
}