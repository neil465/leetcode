func convertToBase7(num int) string {
    res := ""
    if num == 0 {
        return "0"
    }
    b := false
    if num < 0 {
        b = true
        num*=-1
    }
    for num >0 {
        res = fmt.Sprint(num%7) + res
        num/=7
    }
    if b {
        return "-"+res
    }
    return res
    
}