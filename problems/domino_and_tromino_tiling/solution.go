func numTilings(n int) int {
    a := make([]int, n)
    if n <= 2 {
        return n
    }
    a[0],a[1], a[2] = 1, 2, 5
    for i := 3; i < n; i++ {
        a[i] = (a[i - 1] * 2 + a[i - 3]) % 1000000007
    }
    return a[n - 1] 
}