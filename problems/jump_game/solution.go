func canJump(nums []int) bool {
    i := 0
    for dist := 0 ; i <= dist ; i++{
        dist = Max(dist,i+nums[i])
        if i >= len(nums)-1{
            return true
        }
    }
    return false
}
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
