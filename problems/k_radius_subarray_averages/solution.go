func getAverages(nums []int, k int) []int {
    b := 0
    if k * 2  < len(nums)  {
        for i := 0 ; i <= k * 2 - 1; i++ {
            b += nums[i]
        }
    }
    a := make([]int, len(nums))
    for i := 0 ; i < len(nums); i ++ {
        if i < k || i >= len(nums) - k {
            a[i] = -1 
        } else {

            b += nums[i + k]
            a[i] = b / (k * 2 + 1)
            b -= nums[i - k]
            
        }
    }
    return a
}