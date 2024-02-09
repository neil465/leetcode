func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)

    dp := make([][]int, len(nums))
    best := []int{}
    for i := 0; i < len(nums); i ++ {
        dp[i] = []int{nums[i]}
        for j := 0 ; j < i; j ++ {
            if len(dp[j]) >= len(dp[i]) && nums[i] % nums[j] == 0 {
                cpy := make([]int, len(dp[j]))
                copy(cpy, dp[j])
                dp[i] = append(cpy, nums[i])
            }
        }
        if len(dp[i]) > len(best) {
            best = dp[i]
        }
    }

    return best
}