func makeSmallestPalindrome(s string) string {
    a := []rune(s)
    i, j := 0, len(s)- 1 
    for i < j {
        a[i] = min(a[i], a[j])
        a[j] = min(a[i], a[j])
        i ++
        j --
    }
    return string(a)
}
func min(a,b rune) rune {
    if a < b {
        return a
    }
    return b
}