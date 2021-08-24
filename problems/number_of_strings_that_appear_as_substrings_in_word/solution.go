func numOfStrings(patterns []string, word string) int {
    sum := 0
    for _, i := range patterns {if strings.Index(word, i) >= 0 {sum++}}
    return sum
}