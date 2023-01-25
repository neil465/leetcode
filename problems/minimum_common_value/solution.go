func getCommon(nums1 []int, nums2 []int) int {
    var i, j int 

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i ++
        } else if nums1[i] > nums2[j] {
            j ++
        } else {
            return nums1[i]
        }
    }
    return -1
}