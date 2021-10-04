func reverseWords(s string) string {
    a := strings.Fields(s)
    for l,k := range a{
        runes := []rune(k)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        a[l] = string(runes)
    }
    s=strings.Join(a," ")
    return s
}