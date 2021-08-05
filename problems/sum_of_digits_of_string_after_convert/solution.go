var arr []int 
var count int

func getLucky(s string, k int) int {
    
    b := arr
    arr = convert(s)
    
    if len(b) == 0{
        b = arr
    }
    for i := 0 ; i < len(b) ; i++{
        if b[i] < 10{
           continue 
        }
        p := b[i]
        b = append(b[:i], b[i+1:]...)
        t := []int{}
        for p > 0 {
            t = append(t,p%10)
            p /= 10
        }
        b  = append(b[:i], append(t, b[i:]...)...)
        i = -1
    }
    fmt.Println(b)
    res := 0
    for _,i := range b{
        res += i
    }
    a := []int{}
    c := res
    for c > 0 {
        a = append(a,c%10)
        c /= 10
    }
    fmt.Println(res)
    if count == k-1{
        count = 0
        arr = []int{}
        return res 
    }
    count++
    arr = a
    
    return getLucky(s,k)
    
}
func convert (g string) []int{
    res := []int{}
    for _,i := range g{
        res = append(res,int(i-96))
    }
    return res
}