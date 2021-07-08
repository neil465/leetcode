func minProductSum(nums1 []int, nums2 []int) int {
    sum  := 0
    sort.Ints(nums1)
    sort.Slice(nums2,func(i,j int)bool{
        return nums2[i]>nums2[j]
    })
    
    for i,j := range nums1{
        sum += j*nums2[i]
    }
    return sum
    
}