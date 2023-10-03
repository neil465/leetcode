func minOperations(nums []int, k int) int {
    m := map[int]bool{}
    for i := 0; i < k; i++ {
        m[i + 1] = true
    }
    for i := len(nums) - 1; i >= 0 ; i -- {
        if len(m) == 0{
            // fmt.Println(i)
            return len(nums) - i  - 1
        }
        if m[nums[i]] {
            // fmt.Println(m, nums[i])
            delete(m, nums[i])
        }
    }
    return len(nums)
}