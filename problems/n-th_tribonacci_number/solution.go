func tribonacci(n int) int {
    a,b,c := 0,1,1
    if n == 0 {
        return 0
    }
    if n ==1 || n == 2 {
        return 1
    }
    s := 0
    for i:=3;i<=n;i++{
        s = a+b+c
        a,b,c = b,c,s
        
    }
    return s
}