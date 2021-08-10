func isPrefixString(s string, words []string) bool {
    res := ""
    for _,i := range words{
        res +=i
        fmt.Println(res)
        if s == res{
            return true
        }
        
    }
    return false
}