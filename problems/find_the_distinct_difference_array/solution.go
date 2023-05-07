func distinctDifferenceArray(nums []int) []int {
    res := []int{}
    for i := 0 ; i < len(nums) ; i++ {
        prev, end := map[int]bool{}, map[int]bool{}
        for j := 0 ; j <= i; j ++ {
            prev[nums[j]] = true
        }
        for j := i + 1 ; j < len(nums); j++ {
            end[nums[j]] = true
        }

        res = append(res, len(prev) - len(end))
    }
    return res
}