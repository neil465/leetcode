var dp = []int{}

func maximumJumps(nums []int, target int) int {
    dp = make([]int, len(nums))
    for i := range nums {
        dp[i] = math.MinInt32
    }
    b := call(nums, 0, 0, target)
    if b < 0{
        return -1
    }
    return b
}
func call(nums []int, cur int, count int, target int) int {
    if cur >= len(nums) - 1 {
        return 0
    }
    if dp[cur] != math.MinInt32 {
        return dp[cur]
    }
    m := math.MinInt32
    for i := cur + 1; i < len(nums); i ++ {
        if abs(nums[cur] - nums[i]) <= target {
            c := call(nums, i, count + 1, target)
            if c + 1 > m {
                m = c + 1
            }
        }
    }
    dp[cur] = m

    return dp[cur]
}
func abs ( i int ) int {
    if i > 0 {return i}
    return -i
}
func max(i, j int) int {
    if i > j { return i}
    return j
}