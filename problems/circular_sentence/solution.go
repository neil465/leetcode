func isCircularSentence(sentence string) bool {
    a := strings.Fields(sentence)
    a = append(a, a[0])
    for i := 0 ; i < len(a) - 1; i ++ {
        if a[i][len(a[i]) - 1] != a[i + 1][0] {
            return false
        }
    }
    return true
}