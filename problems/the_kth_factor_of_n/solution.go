func kthFactor(n int, k int) int {
    count := 0
    for i := 1; i <= n/2 ; i++{
        if count ==k-1 && n%i == 0{return i}
        if n%i == 0{count++}
    }
    if count== k-1 {return n}
    return -1
}