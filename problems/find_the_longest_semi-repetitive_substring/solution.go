func longestSemiRepetitiveSubstring(s string) int {
    i, j := 0, 0
    c := 0 
    m := 0
    for i < len(s) && j < len(s) {
        if j < len(s) - 1 && s[j] == s[j + 1] {
            if c == 0 {
                j ++
                c++
            } else {
                if s[i] == s[i + 1] {
                    c--
                }
                i ++
            }
        } else if j < len(s) - 1{
            j ++
        } else {
            break
        }
        m = max(m, j - i + 1)
    }
    if len(s) == 1 {
        return 1
    }
    return m 
}
func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}