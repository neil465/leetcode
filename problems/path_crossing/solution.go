func isPathCrossing(path string) bool {
    var dir = map[rune][]int{'N' : {0, 1}, 'S' : {0, -1}, 'E' : {1, 0}, 'W' : {-1 , 0}}

    grid := map[[2]int]bool{{0, 0} : true}

    x, y := 0, 0

    for _, move := range path {
        x, y = x + dir[move][0], y + dir[move][1]
        if grid[[2]int{x, y}] {
            return true
        }
        grid[[2]int{x, y}] = true
    }
    return false
}