func sumOfMultiples(n int) int {
    s := 0
    for i := 1 ; i <= n ; i++ {
        if i % 7 == 0 || i % 5 == 0 || i % 3 == 0 {
            s+= i
        }
    }
    return s
}