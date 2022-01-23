func canJump(nums []int) bool {
    i := 0 
    for dist := 0 ;  i <= dist ; i++{
        dist = max(dist,i+nums[i])
        if i >= len(nums)-1{
            return true
        }
    }
    return false
}
func max (i,j int) int{
    if i > j { return i }
    return j
}
