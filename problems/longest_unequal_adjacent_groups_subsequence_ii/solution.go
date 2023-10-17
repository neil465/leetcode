

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
    m := map[[2]int]bool{}
    for i := range words {
        for j := range words {
            if len(words[i]) == len(words[j]) && hDist(words[i], words[j]) <= 1{
                m[[2]int{i, j}] = true
                m[[2]int{j, i}] = true
            }
        }
    }
    dp = map[int][]string{}
    return f([]string{}, words, groups, 0, -1, m)
}

func hDist(s1, s2 string)int {
    k := 0
    for i := range s1 {
        if s1[i] != s2[i] {
            k ++
        }
    }
    return k
}
var dp = map[int][]string{}


func f(curres []string, words []string, groups []int, curInd, lastInd int , m map[[2]int]bool) []string{
    if _, ok := dp[lastInd]; ok{
        return append(curres, dp[lastInd]...)
    }
    if curInd == len(words) {
        return curres
    }
    k := f([]string{}, words, groups, curInd + 1, lastInd, m)
    
    if lastInd == -1 || (m[[2]int{curInd, lastInd}] && groups[curInd] != groups[lastInd]) {
        v := f([]string{ words[curInd]}, words, groups, curInd + 1, curInd, m)
        if len(v) > len(k) {
            k = v
        }
    }
    dp[lastInd] = k
    return append(curres,dp[lastInd]...)
}