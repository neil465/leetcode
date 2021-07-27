func runningSum(nums []int) []int {
    sum := 0 
    for i,j := range nums{
        sum+=j
        nums[i] = sum
    }
    return nums
}