func threeSum(nums []int) [][]int {
    m := map[int]int{}
    for index, i := range nums {
        m[i] = index
    }
    res := [][]int{}
    used := map[[3]int]bool{}
    for i := range nums {
        for j := i + 1; j < len(nums); j ++ {
            k := !used[[3]int{nums[i] , -(nums[i] + nums[j]), nums[j]}]
            if v, ok := m[-(nums[i] + nums[j])] ; ok && i != v && j != v && k {
                res = append(res, []int{nums[i], nums[j], -(nums[i] + nums[j])})
                used[[3]int{nums[i] , -(nums[i] + nums[j]), nums[j]}] = true
                used[[3]int{nums[j] , -(nums[i] + nums[j]), nums[i]}] = true
                used[[3]int{ -(nums[i] + nums[j]), nums[j], nums[i]}] = true
                used[[3]int{ -(nums[i] + nums[j]), nums[i], nums[j]}] = true
                used[[3]int{ nums[j], nums[i],  -(nums[i] + nums[j])}] = true
                used[[3]int{ nums[i], nums[j],  -(nums[i] + nums[j])}] = true
            }
        }
    }
    return res
}