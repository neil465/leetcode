func rob(nums []int) int {
    memo := make([]int,len(nums))
    for i := 0 ; i < len(memo) ; i++{
        memo[i] = -1
    }
    var robber func([]int,int) int 
    robber = func(nums[]int,pos int) int {
        if pos < 0 {
            return 0
        }
        if memo[pos] > -1 {
            return memo[pos]
        }
        memo[pos] = max(robber(nums,pos-2)+nums[pos],robber(nums,pos-1))
        return memo[pos]
    }
    return robber(nums,len(nums)-1)
}
func max (i,j int) int {
    if i > j {
        return i
    }
    return j
}