func maximumOddBinaryNumber(s string) string {
    k := []rune(s)
    sort.Slice(k, func(i,j int) bool {
        return k[i] > k[j]
    })
    return string(k[1:]) + string(k[0])
}