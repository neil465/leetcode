func generateParenthesis(n int) []string {
    return find(n, n, "")
}
func find(open, closed int, str string) []string {
    if closed == 0 {
        return []string{str}
    }
    
    var res []string
    
    if open < closed {
        res = append(res, find(open, closed - 1, str + ")")...)
    } 
    if open > 0 {
        res = append(res, find(open - 1, closed, str + "(")...)
    }
    
    return res
}