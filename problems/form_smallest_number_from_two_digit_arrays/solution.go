func minNumber(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) && nums1[i] != nums2[j] {
        if nums1[i] < nums2[j]{
            i++
        } else if nums2[j] < nums1[i] {
            j ++
        }
    }   
    if (i >= len(nums1) || j >= len(nums2)) || nums1[i] != nums2[j] {
        return int(math.Min(float64(nums1[0]), float64(nums2[0])) * 10 + math.Max(float64(nums1[0]), float64(nums2[0])))
    }
    return nums1[i]
    
}