
// var dp = map[int]bool{}
func makesquare(matchsticks []int) bool {
    sum := 0
    for _, i := range matchsticks {
        sum += i
    }
    sort.Slice(matchsticks, func(i, j int) bool {
            return matchsticks[i] > matchsticks[j]
        
    })
    if matchsticks[0] > sum / 4 {
        return false
    }
    // dp = map[int]bool{}
    return re(0,0,0,0,0, matchsticks, sum / 4)
}
func re(a,b,c,d, i int, matchsticks []int, t int) bool {
    // if _, ok := dp[i]; ok {
    //     return dp[i]
    // }
    if a > t || b > t || c > t || d> t {
        return false
    }
    if a == b && b == c && c == d && i == len(matchsticks) {
        return true
    } else if i == len(matchsticks) {
        return false
    }
    
    return re(a + matchsticks[i], b, c, d, i + 1, matchsticks, t) || re(a, b + matchsticks[i], c, d, i + 1, matchsticks, t) || re(a, b , c + matchsticks[i], d, i + 1, matchsticks, t) || re(a, b , c , d + matchsticks[i], i + 1, matchsticks, t)
}