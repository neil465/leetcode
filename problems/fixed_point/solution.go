func fixedPoint(arr []int) int {
    res := []int{}
    l,u := 0,len(arr)-1
    for l<=u{
        m := l + (u-l)/2
        if arr[m] == m{
            res = append(res,m)
            u--
            continue
        }
        if arr[m] > m{
            u--
        }else{
            l++
        }
    }
    if len(res) > 0{
        sort.Ints(res)
        fmt.Println(res)
        return res[0]
    }
    return -1
}