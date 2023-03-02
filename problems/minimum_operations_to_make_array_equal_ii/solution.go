func minOperations(nums1 []int, nums2 []int, k int) int64 {
    if k == 0 {
        if reflect.DeepEqual(nums1, nums2) {
            return 0
        } 
        return -1 
    }
    adds := 0
    times := 0
    for i := 0 ; i < len(nums1) ; i++ {
        if abs(nums1[i] - nums2[i]) % k != 0{ 
            return -1
        } else {
            adds += (nums2[i] - nums1[i]) / k
            times += abs(nums2[i] - nums1[i]) / k
        }
    }
    if adds != 0 {
        return -1
    }
    return int64(times/2)
}

func abs(i int) int {
    if i > 0 {
        return i
    }
    return -i
}