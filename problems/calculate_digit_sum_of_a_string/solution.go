func digitSum(s string, k int) string {
    if len(s) <= k {
        return s
    }
    sum := ""
    for i := 0 ; i < len(s) ; i+= k{
        ms := 0
        
        for j := i ; j< int(math.Min(float64(i+k),float64(len(s)))); j++{
            ms += int(s[j])-48
        }
        fmt.Println(ms)
        sum += fmt.Sprint(ms)
    }
    fmt.Println(sum)
    return digitSum(sum,k)
}