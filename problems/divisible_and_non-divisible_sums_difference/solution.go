func differenceOfSums(n int, m int) int {
    n1, n2 := 0,0
    for  i := 1 ; i <= n ; i++ {
        if i % m == 0 {
            n2 += i
        } else {
            n1 += i
        }
    }
    return n1 - n2
}