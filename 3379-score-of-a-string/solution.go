func scoreOfString(s string) int {
    b := 0 
    for i := 1 ;i < len(s); i ++ {
        b += abs(int(s[i]) - int(s[i - 1]))
    }
    return b
}
func abs ( i int) int {
    if i <  0 {
        return -i
    }
    return i
}
