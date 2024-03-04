func longestNiceSubarray(nums []int) int {
    i  := -1
    best := 0
    nums = append(nums, nums[len(nums) -1 ])

    for j := range nums {
        for p := i + 1; p < j ; p ++ {
            if nums[j] & nums[p] != 0 {
                i = p
            }
        }
        best = max(best, j - i)
    }
    // if i == 0 {
    //     return len(nums)
    // }
    return best 
}