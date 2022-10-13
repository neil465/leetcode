func divideString(s string, k int, fill byte) []string {
    i, res := 0, []string{}
    for i = 0 ; i < len(s) - k ; i += k{
        res = append(res, s[i:i + k])
    }
    
    res = append(res, s[i:] + strings.Repeat(string(fill), k - (len(s) - i)))
    
    return res
}