var board [][]int

var directions = [][]int{[]int{0, 1}, []int{0, -1},[]int{1, 0},[]int{-1, 0}, []int{0, 0}}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    board = image
    change( sr, sc, color, map[[2]int]bool{})
    return board
}

func change(y , x, color int, visited map[[2]int]bool) {

    for _, i := range directions {
        if y + i[1] >= 0 && y + i[1] < len(board) && x + i[0] >= 0 && x + i[0] < len(board[0]) && board[y + i[1]][ x + i[0]] == board[y][x] && !visited[[2]int{y + i[1], x + i[0]}]{
                fmt.Println(board)
            visited[[2]int{y + i[1], x + i[0]}]= true
            change(y + i[1], x + i[0], color, visited)
            board[y + i[1]][ x + i[0]] = color
        }
    }
}