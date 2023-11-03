func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    return max(robber(nums[:len(nums) - 1], len(nums), map[int]int{} ), robber(nums[1:], len(nums),  map[int]int{}))
}
func robber(nums []int, pos int, dp map[int]int) int {
    m := 0

    if val, ok := dp[pos]; ok {
        return val
    } 
    for i := 0; i < pos -1 ; i++ {
        m =  max(robber(nums, i, dp) + nums[i] , m) 
    }
    dp[pos] = m
    return m
} 