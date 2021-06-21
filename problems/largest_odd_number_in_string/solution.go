func largestOddNumber(num string) string {
    for i := 0 ; i < len(num) ; i++{
        h, _ := strconv.Atoi(string(num[:len(num)-i][len(num[:len(num)-i])-1]))
        if h%2 != 0 {
            return num[:len(num)-i]
        }
    }
    return ""
}