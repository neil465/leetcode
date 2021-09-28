func searchMatrix(matrix [][]int, target int) bool {
    for j,i := range matrix[0]{
        if i > target{return false}
        for _,i := range matrix{
            if i[j] > target{break}
            if i[j] == target{return true}
        }
    }
    return false
}