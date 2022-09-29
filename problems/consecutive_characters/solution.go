func maxPower(s string) int {
    count, max := 1, 1
    for i := 1 ; i < len(s) ; i++ {
        if s[i] == s[i - 1] {
            count ++
        } else {
            count = 1
        }
        if count > max {
            max = count
        }
    }
    return max
}