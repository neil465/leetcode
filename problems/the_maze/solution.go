var dirs = [][]int{{1,0},{-1,0},{0,1},{0,-1}}

func hasPath(maze [][]int, start []int, destination []int) bool {
    return find(maze, [2]int{start[0], start[1]}, destination, map[[2]int]bool{{start[0],start[1]}: true})
}

func find(maze [][]int, start [2]int, destination []int, pastLoc map[[2]int]bool) bool {
    q := [][2]int{start}
    for len(q) > 0{
        start = q[0]
        if start[0] == destination[0] && start[1] == destination[1] {
            return true
        }
        for _, dir := range dirs {
            tPlace := [2]int{start[0], start[1]}
            for (tPlace[0] >= 0 && tPlace[0] < len(maze) && tPlace[1] >= 0 && tPlace[1] < len(maze[0]) && maze[tPlace[0]][tPlace[1]] != 1){
                tPlace[0] += dir[0]
                tPlace[1] += dir[1]
            }
            tPlace[0] -= dir[0]
            tPlace[1] -= dir[1]
            
            if pastLoc[tPlace] {
                continue
            } else {
                pastLoc[tPlace] = true 

                q = append(q, tPlace)
            }
        }
        q = q[1:]
    }

    return false
    
}