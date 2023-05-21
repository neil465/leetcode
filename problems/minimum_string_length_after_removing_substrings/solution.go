func minLength(s string) int {

    for i := 0 ; i < len(s) - 1; i ++ {

        if i >= 0 {
            if string(s[i: i + 2]) == "AB" || string(s[i: i + 2]) == "CD" {
                s = s[:i] + s[i + 2:]
                i -= 2
            }
        }
    }
    // if len(sa) >=2 && string(sa[:2]) == "AB" || string(sa[:2]) == "CD"  {
    //     sa = sa[2:]
    // }
    // if len(sa) > 2 && string(sa[:len(sa) - 2]) == "AB" || string(sa[:len(sa) - 2]) == "CD"  {
    //     sa = sa[len(sa) - 2:]
    // }
    return len(s)
}