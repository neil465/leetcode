func fib(n int) int {
    a1,a2:= 0 ,1
    if n <2 {return n}
    for i := 1; i < n ; i++{
        sum := a1+a2
        a1 = a2
        a2 = sum
    }
    return a2
}