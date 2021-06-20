func totalMoney(n int) int {
    b := 4

    res := 0
    for i := 0 ; i < n/7 ; i++{
       
        res+= 7*b
        b++
    }
    for i:= 0 ; i< n%7 ; i ++{
        res += i + b-3
    }
    return res
}