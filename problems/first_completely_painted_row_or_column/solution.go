func firstCompleteIndex(arr []int, mat [][]int) int {
    m := map[int][]int{}
     rows := make([]int, len(mat))
    cols := make([]int, len(mat[0]))
    for i := range mat {
        for j := range mat[0] {
            m[mat[i][j]] = []int{i,j}
        }
    }
    for index, v := range arr {
        r, c := m[v][0], m[v][1]
        rows[r] ++
        if rows[r] == len(mat[0]) {
            return index
        }
        cols[c] ++
         if cols[c] == len(mat) {
            return index
        }
    }
    return -1
}