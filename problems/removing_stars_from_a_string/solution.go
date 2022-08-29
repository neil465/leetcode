func removeStars(s string) string {
    a := []rune(s)
    for i := 1 ; i < len(a); i++{
        if a[i] == '*' {
            a = append(a[:i-1], a[i+1:]...)
            i-=2
        }
    }
    if len(a) > 0 && a[0] == '*' {
        a = a[1:]
    }
    return string(a)
}