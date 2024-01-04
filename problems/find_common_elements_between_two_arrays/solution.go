func findIntersectionValues(nums1 []int, nums2 []int) []int {
    m1, m2 := map[int]bool{}, map[int]bool{}

    for _, value := range nums1 {
        m1[value] = true
    }

    one, two := 0,0

    for _, value := range nums2 {
        if m1[value] {
            two ++
        }
        m2[value] = true
    }
    
    for _, value := range nums1 {
        if m2[value] {
            one ++
        }
    }
    return []int{one,two}
}