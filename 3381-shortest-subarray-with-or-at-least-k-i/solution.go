func minimumSubarrayLength(nums []int, k int) int {
    l := math.MaxInt32
    for i := range nums {
        for j := i; j < len(nums); j ++ {
            v := 0 

            for k := i ; k <= j; k ++ {
                v |= nums[k]
            }

            if v >= k && j - i + 1 < l {
                l = j - i + 1
            }
        }
    }
    if l == math.MaxInt32 {
        return -1
    }
    return l
}
