func countTriples(n int) int {
    k := 0
    for i := 1 ; i< n ; i ++ {
        for j := 1 ; j < n; j ++ {
            v := math.Sqrt(float64(i * i + j * j))
           
            if int(v) <= n && float64(int(v)) == v {
                k ++
            }
        }
    }
    return k
}