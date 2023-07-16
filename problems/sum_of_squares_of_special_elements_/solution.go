func sumOfSquares(nums []int) int {
    s := 0
    for i := range nums {
        if len(nums) % (i + 1) == 0{
            fmt.Println(i)
            s += nums[i] * nums[i]
        }
    }
    return s
}