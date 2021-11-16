func checkAlmostEquivalent(word1 string, word2 string) bool {
    chars := make(map[rune]int)   
    for _,i := range word1{
        chars[i]++
    }
    for _,i := range word2{
        chars[i]--
    }
    for _,i := range chars{
        if i > 3 || i < -3{return false}
    }
    return true
}