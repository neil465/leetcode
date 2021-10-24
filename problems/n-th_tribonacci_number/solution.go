func tribonacci(n int) int {
    a,b,c := 0,1,1
    if n == 2{return 1}
    if n <3{return n}
    for i := 2 ; i < n ; i++{
        sum := a+b+c
        a =b 
        b= c
        c = sum
    }
    return c
}