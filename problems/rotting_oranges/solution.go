func orangesRotting(grid [][]int) int {
  
    starting :=[][]int{}
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == 2 {
                starting = append(starting, []int{i,j})
            }
        }
    }
    depth, g := findDepth(grid, starting, 0)
    fmt.Println(g)
    for i := range g {
        for j := range g[0] {
            if grid[i][j] == 1 {
                return -1
            }
        }
    }
    if depth == 0 {
        return 0
    }
    return depth - 1
}

func findDepth(grid [][]int, edgeRO [][]int, depth int) (int, [][]int){
    if len(edgeRO) == 0 {
        return depth, grid
    }
    dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    temp := [][]int{}
    for _, rotted := range edgeRO {
        for _, dir := range dirs {
            if rotted[0] + dir[0] >= 0 && rotted[0] + dir[0] < len(grid) && rotted[1] + dir[1] >= 0 && rotted[1] + dir[1] < len(grid[0]) && grid[rotted[0] + dir[0]][rotted[1] + dir[1]] == 1 {
                temp = append(temp,[]int{rotted[0] + dir[0],rotted[1] + dir[1]})
                grid[rotted[0] + dir[0]][rotted[1] + dir[1]] = 2
            }
        }
    }
    return findDepth(grid, temp, depth+1)
    
}