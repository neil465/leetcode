func incremovableSubarrayCount(nums []int) int {
    val := 0 
    for i := range nums {
        for j := i ; j < len(nums) ; j ++ {
            v := math.MinInt32
            b := true
            for k := range nums {

                if k >= i && k <= j {
                    continue
                }
                if nums[k] > v {
                    v = nums[k]
                } else {

                    b = false
                    break
                }
            }
            if b {
                val ++
            }
        }
    }
    return val
}