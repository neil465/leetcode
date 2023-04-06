func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    m1, m2 := map[int]bool{},  map[int]bool{}
    res := []int{}
    contained := map[int]bool{}
    for _, i := range nums1 {
        m1[i] = true
    }
    for _, i := range nums2 {
        if m1[i] && !contained[i]{
            res = append(res, i)
            contained[i] = true
        }
        m2[i] = true
    }
    for _, i := range nums3{
        if (m1[i] || m2[i]) && !contained[i]{
            res = append(res, i)
            contained[i] = true
        }
    }
    return res
}