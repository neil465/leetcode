func reverseWords(s string) string {
    a := strings.Fields(s)
    sum := ""
    for i := len(a) - 1 ; i >= 0 ; i-- {
        sum += a[i] + " "
    }
    return sum[:len(sum) - 1]
}