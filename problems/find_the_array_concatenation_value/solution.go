func findTheArrayConcVal(nums []int) int64 {
    sum := 0
    for len(nums) > 0 {
        first := strconv.Itoa(nums[0])
        if len(nums) > 1 {
            second := strconv.Itoa(nums[len(nums) - 1])
            concat, _ := strconv.Atoi(first + second)

            sum += concat
            nums = nums[1:len(nums) - 1]

        } else {
            sum += nums[0]
            nums = nums[1:]
        }
    }
    return int64(sum)
}