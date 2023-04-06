func robotSim(commands []int, obstacles [][]int) int {
    directions := [][]int{{0 , 1}, {1, 0}, {0, -1}, {-1, 0}}
    dir := 0
    curX, curY := 0,0 
    mDist := 0
    obstacleMap := map[[2]int]bool{}

    for _, i := range obstacles {
        obstacleMap[[2]int{i[0], i[1]}] = true
    }

    for _, i := range commands {
        if i == -2 {
            dir --
            if dir == -1 { dir = 3 }
        } else if i == -1 {
            dir ++
            if dir ==  4 { dir = 0 }
        } else {
            for j := 0 ; j < i ; j++ {
                
                if obstacleMap[[2]int{curX + directions[dir][0], curY + directions[dir][1]}] {
                    
                    continue
                } 
                curX += directions[dir][0]
                curY += directions[dir][1]
            }
        }
        if mDist < curX * curX + curY * curY {
            mDist = curX * curX + curY * curY
        }
    }
    return mDist
}