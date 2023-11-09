func canWin(currentState string) bool {
    return a([]byte(currentState), map[string]bool{})
}
func a(cstate []byte, dp map[string]bool) bool {
    if val, ok := dp[string(cstate)] ; ok {
        return val
    }
    for i := 0 ; i < len(cstate) -1 ; i ++ {
        if string(cstate[i]) + string(cstate[i + 1]) == "++" {
            cstate[i], cstate[i + 1] = byte('-'), byte('-')
            v := !a(cstate, dp)
            cstate[i], cstate[i + 1] = byte('+'), byte('+')
            if v {
                dp[string(cstate)] = true
                return true
            }
        }
    }
    dp[string(cstate)] = false
    return false
}