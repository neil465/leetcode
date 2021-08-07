func getNoZeroIntegers(n int) []int {
    min, max := 1,n-1
    for true{
        mincopy := min
        maxcopy := max
        f,f2 := true,true
        for mincopy > 0{
            if mincopy%10 == 0{
                f= false
            }
            mincopy /= 10
        }
        for maxcopy > 0{
            if maxcopy%10 == 0{
                f= false
            }
            maxcopy /= 10
        }
        if f == true && f2 == true{
            return []int{min,max}
        }
        min++
        max--
    }
    return []int{}
}