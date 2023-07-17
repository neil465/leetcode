func maximumBeauty(nums []int, k int) int {
    sort.Ints(nums)
    start := 0 
    max := 0
    for i := range nums {
        for nums[i] - nums[start] > 2 * k {
            start ++
        }
        if i - start + 1 > max {
            max = i - start +1
        }
    }
    return max
}
