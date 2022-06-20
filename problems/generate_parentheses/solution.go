func generateParenthesis(n int) []string {
    return backtrack(n,0,0,"")
}
func backtrack(m, o, c int, s string ) []string{
    if len(s) >= m * 2 {
        return []string{s}
    }
    temp := []string{}
    if o < m {
        temp = append( temp, backtrack(m, o + 1, c, s + "(" )...)
    }
    if c < o {
        temp = append( temp, backtrack(m, o , c + 1, s + ")" )...)
    }
    return temp
}