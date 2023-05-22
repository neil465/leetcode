type pos struct {
    row, col, dist int
}

func nearestExit(maze [][]byte, entrance []int) int {
    q := []pos{{entrance[0], entrance[1], 0}}
    p := pos{}
    dirs := []pos{{0,1,0},{0,-1,0},{1,0,0},{-1,0,0}}
    for len(q) > 0 {
        q, p = q[1:], q[0]
        if (p.row == 0 || p.row == len(maze) - 1|| p.col == 0 || p.col == len(maze[0]) - 1) && (p.row != entrance[0] || p.col != entrance[1]){
            return p.dist
        }
        for _, dir := range dirs {
            tr, tc := p.row + dir.row, p.col + dir.col
            if tr >= 0 && tr < len(maze) && tc >= 0 && tc < len(maze[0]) && maze[tr][tc] != '+'{
                maze[tr][tc] = '+'
                q = append(q, pos{tr ,tc ,p.dist + 1} )
            }

        }

    }
    return -1
}