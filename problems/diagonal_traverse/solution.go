func findDiagonalOrder(mat [][]int) []int {
    m := [100000][]int{}
    for i := range mat {
        for j := range mat[i] {
            m[i + j] = append([]int{mat[i][j]}, m[i + j]...)
        }
    }
    a := []int{}
    for index, i := range m {
        if index % 2 != 0 {
            t := []int{}
            for _, v:= range i {
                t = append([]int{v}, t...)
            }
            a = append(a, t...)
        } else {
            a = append(a, i...)
        }
        
    }
    return a
}