func maximumNumberOfStringPairs(words []string) int {
    b := 0
    for i := range words {
        for j :=  i + 1; j < len(words); j ++ {
            if words[i][0] == words[j][1] && words[i][1] == words[j][0] {
                b++
            }
        }
    }
    return b
}