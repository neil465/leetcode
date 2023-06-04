func minimizedStringLength(s string) int {
    m := map[rune]bool{}
    for _, i := range s {
        m[i] = true
    }
    return len(m)
}