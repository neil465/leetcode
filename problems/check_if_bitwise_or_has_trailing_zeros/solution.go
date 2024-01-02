func hasTrailingZeros(nums []int) bool {
    for i := range nums {
        for j := range nums {
            if i != j && nums[i] & 1 == 0 && nums[j] & 1 == 0 {
                return true
            }
        }
    }
    return false
}