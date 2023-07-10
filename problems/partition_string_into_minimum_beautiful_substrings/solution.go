func minimumBeautifulSubstrings(s string) int {
    b := helper(s) 
    if b == math.MaxInt32 {
        return -1
    }
    return b
}
var m = map[string]bool{"11110100001001" : true, "110000110101" : true, "1001110001" : true, "1111101" : true, "11001" : true, "101" : true, "1" : true}
func helper(s string) int  {
    fmt.Println(s)
    if s == "" {
        return 0
    }
    min := math.MaxInt32
    for i := 0 ; i < len(s); i ++ {
        if m[s[:i + 1]] {
            a  := helper(s[i + 1:]) + 1
            if a < min {
                min = a
            }
        }
    }
    return min
}
