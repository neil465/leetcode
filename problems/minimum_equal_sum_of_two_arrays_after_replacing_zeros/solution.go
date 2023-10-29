func minSum(nums1 []int, nums2 []int) int64 {
    cursum1, cursum2 := 0, 0
    c1, c2 := 0,0
    for _, i := range nums1 {
        cursum1 += i
        if i == 0 {
            c1 ++
        }
    }
    for _, i := range nums2 {
        cursum2 += i
        if i == 0 {
            c2 ++
        }
    }
    
    if cursum1 + c1 >= cursum2 + c2 && c2 != 0 {
        return int64(cursum1 + c1)
    } else if  cursum1 + c1 <= cursum2 + c2 && c1 != 0 {
        return int64(cursum2 + c2)
    } else if cursum1 + c1 == cursum2 + c2 {
        return int64(cursum1 + c1 )
    }
    return -1
    
    
    
}