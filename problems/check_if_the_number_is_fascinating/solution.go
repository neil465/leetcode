func isFascinating(n int) bool {
    a, b := n * 2, n* 3 
    m := map[int]int{}
    for n > 0 {
        m[n % 10] ++
        n /= 10
    }
    for b > 0 {
        m[b % 10] ++
        b /= 10
    }
    for a > 0 {
        m[a % 10] ++
        a /= 10
    }
    if len(m) != 9 {
        return false
    }
    if m[0] > 0 {
        return false
    }
    for i := 1 ; i < 10; i++ {
        if m[i] != 1 {
            return false
        }
    }
    return true
    
}