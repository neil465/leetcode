func minOperations(nums []int, k int) int {
    sort.Ints(nums) 

    for i := 0 ; i < len(nums); i ++ {
        if nums[i] >= k {
            return i
        }
    }
    return len(nums) - 1
}