func confusingNumber(n int) bool {
    if n == 0 {
        return false
    }
    m := map[int]int{ 0 : 0, 1 : 1, 6 : 9, 8 : 8, 9 : 6}
    res := ""
    sn := strconv.Itoa(n)
    for n > 0 {
        a := n % 10

        n /= 10
        if _, ok := m[a]; !ok {
            return false
        }
        res += string(48 + m[a])
    }
    return res != sn
}