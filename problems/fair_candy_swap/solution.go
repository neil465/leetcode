func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    result := 0
    result2 := 0
    m := make(map[int]int)
    for _, v := range bobSizes {  
        result += v  
    }  
    for _, v := range aliceSizes {  
        result2 += v  
        m[v] = 2
    } 
    target := (result+result2)/2
    for _,i := range bobSizes{
        if m[result2+i-target]==2{
            return []int{result2+i-target,i}
        }
    } 
    panic("unreachable")
}