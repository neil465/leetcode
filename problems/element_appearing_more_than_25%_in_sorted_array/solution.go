func findSpecialInteger(arr []int) int {
    if len(arr)<=2{
        return arr[0]
    }
    c,x,v := 1,0,0
    for i,_ := range arr{
        if i == len(arr)-1{
            if arr[i] == arr[i-1]{
                c++
            }
            if c >= x{
                x = c
                v = arr[i]
                return v
            }
            continue
        }
        if arr[i] == arr[i+1]{
            c++
        }else if c >= x {
            v =arr[i]
            x =c
            c = 1
        }
    }
    fmt.Println(c)
    return v
}