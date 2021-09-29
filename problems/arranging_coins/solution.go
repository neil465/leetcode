func arrangeCoins(n int) int {
    count := 0
    i := 1
    for n >= 0{
        n-=i
        if n >= 0{count++}
        i++
    }
    return count
}