func maximumTripletValue(nums []int) int64 {
    max := 0
    for i := range nums {for j := i + 1 ; j < len(nums); j ++ {for k := j + 1 ; k < len(nums); k ++ {if (nums[i] - nums[j]) * nums[k] > max {max =(nums[i] - nums[j]) * nums[k]}}}}
    return int64(max)
}