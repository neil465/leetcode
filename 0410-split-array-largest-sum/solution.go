func splitArray(nums []int, k int) int {
    dp := make([][]int, k)
    dp[0] = make([]int, len(nums))
    cursum := 0 

    for i := range nums {
        cursum += nums[i]
        dp[0][i] = cursum
    }

    for curk := 1 ; curk < k ; curk ++ {
        dp[curk] = make([]int, len(nums))
        for i := curk; i < len(nums); i ++ {
            dp[curk][i] = math.MaxInt32
            for j := curk; j < len(nums); j ++ {
                dp[curk][i] = min(dp[curk][i], max(dp[curk - 1][j - 1], dp[0][i] - dp[0][j - 1 ]))
            }
        }
    }

    return dp[len(dp) - 1][len(dp[0]) - 1]
}
