var res = []string{}

func findWords(board [][]byte, words []string) []string {
    m := map[string]bool{}
    res = []string{}
    wordIsh := map[string]bool{}
    for _, i := range words {
        for j := 1; j <= len(i); j++ {
            wordIsh[i[:j]] = true
        }
        m[i] = true
    }
    for i := range board {
        for j := range board[0] {
            helper(wordIsh, i, j, m, map[[2]int]bool{[2]int{i, j} : true}, string(board[i][j]), board)
        }
    }
    return res
}
func helper(wordIsh map[string]bool, i, j int, m map[string]bool, visited map[[2]int]bool, cur string, board [][]byte) {
    if !wordIsh[cur] {
        return
    }
    if m[cur] {
        
        res = append(res, cur)
        m[cur] = false
    }   
    if len(cur) ==11 {
        return 
    }
    
    
    for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
        if i + dir[0] >= 0 && i + dir[0] < len(board) && j + dir[1] >= 0 && j + dir[1] < len(board[0]) && visited[[2]int{i + dir[0], j + dir[1]}] == false {
            visited[[2]int{i + dir[0], j + dir[1]}] = true
            helper(wordIsh, i + dir[0], j + dir[1], m, visited, cur + string(board[i + dir[0]][j + dir[1]]), board )
            visited[[2]int{i + dir[0], j + dir[1]}] = false
        }
    }
}