func intersection(nums1 []int, nums2 []int) []int {
    
    res := []int{}
    mappy := make(map[int]int)
    for _,i := range nums1{
        mappy[i]++
    }
    for _,i := range nums2{
        if mappy[i] > 0 {
            res = append(res,i)
            mappy[i] = 0
        }
    }
    return res
    

}