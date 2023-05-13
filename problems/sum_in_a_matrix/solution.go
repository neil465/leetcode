func matrixSum(nums [][]int) int {
    for k := range nums {
        sort.Slice(nums[k], func(i,j int)bool {
            return nums[k][i] > nums[k][j]
        })
    }
    a := 0

    for len(nums[0]) > 0 {
        h := 0
        for i := range nums {
            if nums[i][0] > h {
                h = nums[i][0]
            }
            nums[i] = nums[i][1:]
        }
        a += h
    }
    return a
}