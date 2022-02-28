var dp [100000]int

func rob(nums []int) int {
    dp = [100000]int{}
    return robber(nums,len(nums)-1)
}
func robber(nums []int, pos int) int {
    if pos < 0 {
        return 0
    }
    if dp[pos] <= 0 {
        dp[pos] = int(math.Max(float64(robber(nums,pos-2)+nums[pos]),float64(robber(nums,pos-1))))+1    
    }
    return dp[pos]-1
}
