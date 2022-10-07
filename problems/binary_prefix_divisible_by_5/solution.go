func prefixesDivBy5(nums []int) []bool {
    result, sum := make([]bool , len(nums)), 0
    for i := 0 ; i < len(nums) ; i++{
        sum = (sum * 2 + nums[i]) % 5
        result[i] = sum == 0
    }
    return result
}