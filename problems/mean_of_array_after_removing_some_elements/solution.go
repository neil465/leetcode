func trimMean(arr []int) float64 {
    sort.Ints(arr)
    b := len(arr) / 20 
    a := 0
    for i := b ; i < len(arr) - b; i++{
        a += arr[i]
    }
    return float64(a) / float64(len(arr) - (b * 2))
}