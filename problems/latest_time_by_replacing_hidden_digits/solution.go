func maximumTime(time string) string {
    a := []rune(time)
    if a[0] == '?' {
        
        a[0] = '2'
        if a[1] != '?' && a[1] > '3' {
            a[0] = '1'
        }
    }
    if a[1] == '?' {
        a[1] = '3'
        if a[0] < '2' {
            a[1] = '9'
        }
    }
    if a[3] == '?' {
        a[3] = '5'
    }
    if a[4] == '?' {
        a[4] = '9'
    }
    return string(a)
}