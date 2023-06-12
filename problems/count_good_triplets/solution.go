func countGoodTriplets(arr []int, a int, b int, c int) int {
    g := 0
    for i := range arr {
        for j := i + 1 ; j < len(arr); j++ {
            for k := j + 1 ; k < len(arr); k++ {
                if abs(arr[i] - arr[j]) <= a && abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c {
                    g ++
                }
            }
        }
    }
    return g
}
func abs(i int) int {
    if i > 0 { return i }
    return i * -1 
}