
var board = []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}

func alphabetBoardPath(target string) string {
    sum := ""
    currow, curcol := 0, 0
    for _, i := range target {
        newRow, newCol := int(i - 'a') / 5, int(i - 'a') % 5 
        fmt.Println(currow, curcol, newRow, newCol)
        res, _, _ := find(currow, curcol, newRow, newCol, "")
        sum += res + "!"
        currow, curcol = newRow, newCol
    }
    return sum
}
func find(row, col, row2, col2 int, sum string) (string, int, int) {
    if row == row2 && col == col2{
        return sum, row,col
    }
    if board[row][col] == byte('z') || row2 < row{
        row --
        sum += "U"
        return find(row, col, row2, col2, sum)
    }
    if row2 > row && (row != 4 || col == 0){
        row ++
        sum += "D"
        return find(row, col, row2, col2, sum)
    } 
    if col2 > col{
        col ++
        sum += "R"
        return find(row, col, row2, col2, sum)
    } 
     if col2 < col{
        col --
        sum += "L"
        return find(row, col, row2, col2, sum)
    } 
    return sum, row, col

}