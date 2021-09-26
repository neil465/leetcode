func maximumDifference(nums []int) int {
    max := -1
    for i := range nums{
        for j := i+1 ; j < len(nums) ; j++{
            if nums[i] < nums[j] && nums[j]-nums[i] > max{
                max = nums[j]-nums[i]
            }
        }
    }
    return max
}