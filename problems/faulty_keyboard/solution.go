func finalString(s string) string {
    a := []rune{}
    for _, v := range s {
        if v == 'i' {
            for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
                a[i], a[j] = a[j], a[i]
            }
        } else {
            a = append(a, v)
        }
    }
    return string(a)
}