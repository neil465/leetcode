func rotateString(s string, goal string) bool {
    if len(s) == 0 && len(goal) == 0{
        return true
    }
    for i := 0 ; i < len(s) ; i++{
        if s == goal{
            return true
        }
        s = string(s[len(s)-1]) + s[:len(s)-1]
        
    }
    return false
}