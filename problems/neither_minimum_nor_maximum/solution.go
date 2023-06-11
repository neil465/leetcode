func findNonMinOrMax(nums []int) int {
    max, min := 0, 10000
    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] > max {
            max = nums[i]
        }
        if nums[i] < min {
            min = nums[i]
        }
    }
    for i := range nums {
        if nums[i] != max && nums[i] != min {
            return nums[i]
        }
    }
    return -1
}