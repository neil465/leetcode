func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    min := 100000000000000
    for k,i := range wordsDict{
        if i == word1{
            for n,j := range wordsDict{
                if j == word2 && min > abs(n-k){
                    min = abs(n-k)
                }
            }
        }
    }
    return min
}
func abs (i int)int {
    if i < 0 {
        return i*-1
    }
    return i
}