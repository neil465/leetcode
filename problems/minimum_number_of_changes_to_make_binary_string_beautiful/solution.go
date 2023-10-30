func minChanges(s string) int {
    c := 0 
    t := false
    if s[0] == '1' {
        t = true
    }
    res := 0 
    for i := range s {
        if (s[i] == '1') == t {
            c ++
        } else {
            if c % 2 != 0 {
                c ++
                res ++
            } else {
                c= 1
                t = !t
            }
        }
    }
    return res
}