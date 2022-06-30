func minMoves2(nums []int) int {
    sort.Ints(nums)
    mid := nums[len(nums) / 2]
    count := 0 
    for i := 0 ; i < len(nums) ; i++{
        count += int(math.Abs(float64(nums[i] - mid)))
    }
    return count
}