func anagramMappings(nums1 []int, nums2 []int) []int {
    arr := []int{}
    for _,i := range nums1{
        for j,k := range nums2{
            if k == i {
                arr = append(arr,j)
                break
            }
        }
    }
    return arr
}