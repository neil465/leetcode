func sumDistance(nums []int, s string, d int) int {
    for i := range nums {
        if s[i] == 'R' {
            nums[i ] += d
        } else {
            nums[i] -= d
        }
    }
    dist := 0
    sort.Ints(nums)
    rsum := 0
    for i := range nums {

        dist += (nums[i] * i - rsum) 
        dist %= 1000000007
        rsum += (nums[i]) 
        rsum %= 1000000007
    }
    return dist % 1000000007

}
func abs(i int) int {
    if i < 0 {
        return i * -1
    }
    return i
}