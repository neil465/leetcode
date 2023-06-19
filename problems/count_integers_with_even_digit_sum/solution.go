func countEven(num int) int {
    g := 0
    for i := 1 ; i <= num; i ++ {
        a := i
        s := 0
        for a > 0 {
            s += a % 10
            a /= 10
        }
        if s % 2== 0  {

            g ++
        }
    }
    return g
}