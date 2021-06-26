func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    min := 999999999909099090
    w1 := -1
    w2 := -1
    
    for i := 0 ; i < len(wordsDict) ; i++{
        if wordsDict[i] == word1{
            w1 = i
            if w2 != -1 && abs(w2 - w1) < min{
                min = abs(w2 - w1)
            }
        }
        if wordsDict[i] == word2{
            w2 = i
            if w1 != -1 && abs(w1 - w2) < min{
                min = abs(w1 - w2)
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