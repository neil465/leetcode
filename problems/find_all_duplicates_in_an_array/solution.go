func findDuplicates(nums []int) []int {
    sort.Ints(nums)
    a := []int{}
    for i := 0 ; i < len(nums) - 1; i ++ {
        if nums[i] == nums[i + 1] {
            a = append(a, nums[i])
        }
    }
    return a
}