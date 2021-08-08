func thousandSeparator(n int) string {
    arr,flag,count,res := fmt.Sprint(n) ,false,0,""
    
    if n < 1000{return fmt.Sprint(n)}
    if len(arr)%3 == 0{flag = true}
    
    for i := len(arr)-1 ; i >= 0 ; i --{
        count++
        res = string(arr[i]) + res
        if count == 3{
            res = "." + res
            count = 0
        }
    }
    
    if flag{return res[1:]}
    
    return res
}