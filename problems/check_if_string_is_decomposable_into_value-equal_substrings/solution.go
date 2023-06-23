func isDecomposable(s string) bool {
    b := false
    if len(s) ==2 && s[0] == s[1]{ return true}
    if len(s) <= 3{
        return false
    }
    for i := 1 ; i < len(s); i += 3 {
        if ((i >= len(s) - 2) || s[i] != s[i + 1]) && s[i] != s[i - 1]{
            return false
            
        }
        fmt.Println(s)
            lc := []rune(strings.Repeat(" ", len(s)))
            lc[i] = rune(s[i])
            fmt.Println(string(lc))
        if (i >= len(s) - 1) {
            if s[i] != s[i - 1] {
                if b {
                    fmt.Println("K2")
                    return false
                }   
                
            }
            b = true
        } else if s[i] != s[i + 1] || s[i] != s[i - 1]{
            if b {
                fmt.Println("K")

                return false
            }
            
            b = true
            i--
        }

        fmt.Println(len(s) - 1 - i, len(s) -1)
    }
    
    if !b {
        fmt.Println("hi")
        return false
    }
    return true
}