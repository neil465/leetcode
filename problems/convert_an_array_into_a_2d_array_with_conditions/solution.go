func findMatrix(nums []int) [][]int {
    m := map[int]int{}
    for _, val := range nums {
        m[val] ++
    }
    res := [][]int{}
    for val, count := range m {
        for ind := 0; ind < count; ind ++ {
            if ind == len(res) {
                res = append(res, []int{val})
            } else {
                res[ind] = append(res[ind], val)
            }
        }
    }
    return res
}

func setBit(k, n int) int {
    return ((1 << k) | n)
}
func unsetBit(k, n int) int {
    return (^(1 << k) & n)
}
func isSet(k, n int) bool {
    return (1 << k) & n > 0
}