func arithmeticTriplets(nums []int, diff int) int {
    sum, m := 0, [1000]bool{}
    for _, i := range nums {
        if i >= 2 * diff && m[i - diff] && m[i - 2 * diff] {
            sum ++
        }
        m[i] = true
    }
    return sum
}