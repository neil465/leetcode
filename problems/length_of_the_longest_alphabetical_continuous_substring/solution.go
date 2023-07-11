func longestContinuousSubstring(s string) int {
    c, mc := 1, 1
    for i := 1; i < len(s); i ++ {
        if s[i] - s[i - 1] != 1 {
            c = 1
        } else {
            c ++
            if c > mc {
                mc = c
            }
        }
    }
    return mc
}