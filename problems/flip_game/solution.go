func generatePossibleNextMoves(currentState string) []string {
    a := []byte(currentState)
    res := []string{}

    for i := 0 ; i < len(a) - 1 ; i++ {
        if a[i] == byte('+') && a[i+1] == byte('+') {
            a[i], a[i+1] = byte('-'), byte('-')
            res = append(res, string(a))
            a[i], a[i+1] = byte('+'), byte('+')
        }
    }
    return res
}