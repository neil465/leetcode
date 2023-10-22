func minimumSum(nums []int) int {
    minsum := math.MaxInt32
    for i := range nums {
        for j := i + 1; j < len(nums); j ++ {
            for k := j + 1; k < len(nums); k ++ {
                if nums[j] > nums[i] && nums[j] > nums[k] {
                    if nums[i] + nums[j]+ nums[k] < minsum {
                        minsum = nums[i] + nums[j]+ nums[k]
                    }
                }
            }
        }
    }
    if minsum == math.MaxInt32 {
        return -1
    }
    return minsum
}