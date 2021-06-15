func xorOperation(n int, start int) int {
    arr := []int{}
    res := start
    for i := 0 ; i < n ; i++{
        arr = append(arr,(i*2)+start)
    }
    fmt.Println(arr)
    for i := 1 ; i < len(arr) ; i++{
        res ^= arr[i]
            fmt.Println(res)
}
    return res
}