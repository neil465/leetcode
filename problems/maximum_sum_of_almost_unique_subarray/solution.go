func maxSum(nums []int, m int, k int) int64 {
    ma := map[int]int{}
    sum := 0 
    maxSum := 0
    for i := 0; i < k ; i ++ {
        ma[nums[i]] ++
        sum += nums[i]
    }
    if len(ma) >= m {
        maxSum = sum
    }
    
    for i := k  ; i < len(nums); i++ {
        sum += nums[i] - nums[i - k]
        ma[nums[i - k]] --
        ma[nums[i]] ++
        if ma[nums[i - k]] == 0 {
            delete(ma, nums[i - k])
        }
        if len(ma) >= m && sum > maxSum {
            maxSum = sum 
        }
    }
    
    return int64(maxSum)
}