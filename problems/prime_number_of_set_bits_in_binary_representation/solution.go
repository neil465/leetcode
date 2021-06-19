func countPrimeSetBits(left int, right int) int {
    res := 0
    for i := left ; i <= right ; i++{
        minires := 0
        b := i
        for b >0{
            if b&1 == 1{
                minires++
            }
            b = b>>1
        }
        flag := true
        fmt.Println(minires)
        if minires == 2{
            fmt.Println(i)
            res++
            continue
        }
        if minires == 1{
            flag = false
        }
        for i := 2 ; i < minires;i++{
            if minires%i == 0{
                flag = false
                break
            }
        }
        if flag {
            fmt.Println(i)
            res++
        }
    }
    return res
    
}