func arrayRankTransform(arr []int) []int {
    b := make([]int,len(arr))
    copy(b,arr)
    res := make([]int,len(arr))
    counter := 0
    mappy := make(map[int]int)
    sort.Ints(b)
    for _,i2 := range b{
        
        if _, ok := mappy[i2]; ok {
            continue
        }else{
         counter++
            mappy[i2] = counter
        }
    }
    for i :=0 ; i<len(arr) ; i++{
        res[i] = mappy[arr[i]]
    }
    return res
}