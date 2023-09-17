func sumIndicesWithKSetBits(nums []int, k int) int {
    s := 0
    for i := range nums {
        sum := 0
        t := i
        for t > 0 {
            sum += t & 1
            t /= 2
        }
        if sum == k {
            s += nums[i]
        }
    }
    return s
}