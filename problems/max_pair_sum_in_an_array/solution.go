func maxSum(nums []int) int {
    max := -1
    sort.Ints(nums)
    for v := 0; v < 10; v++ {
        k := []int{}
        for i := len(nums) -1 ; i >= 0 ; i -- {
            value := nums[i]
            max := 0 
            for value > 0 {
                if value % 10 > max {
                    max = value % 10
                }
                value /= 10
            }
            if max == v {
                k = append(k, nums[i])
            }
        }

        if len(k) >= 2 && k[1] + k[0] > max {
            max = k[1] + k[0]
        }
    }
    return max
}