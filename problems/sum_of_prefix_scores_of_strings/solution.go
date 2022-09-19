func sumPrefixScores(words []string) []int {
    m, result := map[string]int{}, []int{}
    for i := 0; i < len(words); i++ {
        for j := 0 ; j < len(words[i]); j++{
            m[words[i][:j+1]]++
        }
    }
    for i := 0; i < len(words); i++ {
        sum := 0
        for j := 0 ; j < len(words[i]); j++{
            sum += m[words[i][:j+1]]
        }
        result = append(result, sum)
    }
    return result
    
}