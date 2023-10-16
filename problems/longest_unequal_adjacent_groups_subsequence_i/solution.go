var dp = map[int][]string{}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
    dp = map[int][]string{}
    return f([]string{}, words, groups, 0, -1)
}
func f(curres []string, words []string, groups []int, curInd, lastInd int ) []string{
    if _, ok := dp[lastInd]; ok{
        return append(curres, dp[lastInd]...)
    }
    if curInd == len(words) {
        return curres
    }
    k := f([]string{}, words, groups, curInd + 1, lastInd)
    
    if lastInd == -1 || groups[curInd] != groups[lastInd] {
        v := f([]string{ words[curInd]}, words, groups, curInd + 1, curInd)
        if len(v) > len(k) {
            k = v
        }
    }
    dp[lastInd] = k
    return append(curres,dp[lastInd]...)
}