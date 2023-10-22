func minimumSum(nums []int) int {
    preMin := make([]int, len(nums))
    postMin := make([]int, len(nums))
    preMin[0] = nums[0]
    postMin[len(nums) - 1] = nums[len(nums) - 1]
    for i := 1 ; i< len(nums); i ++ {
        preMin[i] = min(preMin[i - 1], nums[i - 1])
    }
    for i := len(nums) -2 ; i >= 0 ; i -- {
        postMin[i] = min(postMin[i + 1], nums[i + 1])
    }
    min := math.MaxInt32
    for i := 1; i < len(nums) - 1; i ++ {
        if nums[i] > preMin[i] && nums[i] > postMin[i] && postMin[i] + preMin[i] + nums[i] < min {
            min = postMin[i] + preMin[i] + nums[i]
        } 
    }
    if min == math.MaxInt32 {
        return - 1
    }
    return min
}
func min(i, j int)int {
    if i < j {
        return i
    }
    return j
}