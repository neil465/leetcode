func lengthOfLIS(nums []int) int {
    curL := 0
    dp := make([]int, len(nums))

    for i := len(nums) -1 ; i >= 0 ; i -- {
        for j := i + 1 ; j < len(nums); j++ {
            if nums[j] > nums[i] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        } 
        curL = max(curL, dp[i])
    }
    return curL + 1
}