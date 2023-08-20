func minimumSum(n int, k int) int {
    m := map[int]bool{}
    s := 0 
    sum := 0 
    for i := 1; s < n; i++ {
        if !m[i] {
            s ++
            sum += i
        }
        m[k - i] = true
    }
    return sum
    
}