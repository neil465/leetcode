var dirs = [][]int{{0,1},{0,-1},{1,0},{-1,0}}

func exist(board [][]byte, word string) bool {
    for i := range board {
        for j := range board[i] {
            if board[i][j] == word[0] && find(board, word[1:], i, j, map[[2]int]bool{[2]int{i, j} : true}) {
                return true
            }
        }
    }
    return false
    
}

func find(board [][]byte, word string, i, j int, visited map[[2]int]bool) bool {
    if word == "" {
        return true
    }
    for _, dir := range dirs {
        if i + dir[0] >= 0 && i + dir[0] < len(board) && j + dir[1] >= 0 && j + dir[1] < len(board[0]) && board[i + dir[0]][j + dir[1]] == word[0] && !visited[[2]int{i + dir[0], j + dir[1]}] {
            visited[[2]int{i + dir[0], j + dir[1]}] = true
            if find(board, word[1:], i + dir[0], j + dir[1], visited) {
                return true
            }
            visited[[2]int{i + dir[0], j + dir[1]}] = false
        }
    }
    return false
}
