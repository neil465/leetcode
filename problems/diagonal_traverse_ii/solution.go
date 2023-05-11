func findDiagonalOrder(nums [][]int) []int {
    m := [100000][]int{}
    for i := range nums {
        for j := range nums[i] {
            m[i + j] = append([]int{nums[i][j]}, m[i + j]...)
        }
    }
    a := []int{}
    for _, i := range m {
        a = append(a, i...)
    }
    return a
}

//[14,12,19,16, 9],
//[13,14,15, 8,11],
//[11,13, 1]]