func largestSubarray(nums []int, k int) []int {
    max := 0
    maxIndex := 0
    for i := 0 ; i < len(nums)-k+1 ; i++{
        if nums[i] > max{
            max = nums[i]
            maxIndex = i
        }
    }
    return nums[maxIndex:maxIndex+k]
}