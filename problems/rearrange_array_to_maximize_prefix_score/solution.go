func maxScore(nums []int) int {
    sort.Ints(nums)

    sum := 0
    count := 0
    for i := len(nums) - 1; i >= 0; i--{
        sum += nums[i]

        if sum > 0 {
            count++
        }
    }

    return count
}