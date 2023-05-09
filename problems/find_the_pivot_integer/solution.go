func pivotInteger(n int) int {
    s := summation(n)
    for i := 0 ; i < n ; i ++ {
        if s - summation(i) == summation(i + 1) {
            return i + 1
        }
    }
    return -1

}
func summation(n int) int {
    return (n * (n + 1))/2
}