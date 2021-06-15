func findComplement(num int) int {
    count := 0 
    b := num
    for b >0 {
        b=b>>1
        count++
    }
    res := 1
    for i := 1 ; i < count ; i++{
        res*=2
        res++
    }
    return num^res
}