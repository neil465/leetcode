func maxAreaOfIsland(grid [][]int) int {
    max := 0
    for i := 0 ; i < len(grid) ; i++{
        for j := 0 ; j < len(grid[0]); j++{
            if grid[i][j] == 1{
                a,b := countandfill(grid,i,j,0)
                if a > max{max = a}
                grid = b
            }
        }
    }
    
    return max
}

type key struct {
    y int
    x int
}

func countandfill(image [][]int, sr int, sc int, newColor int) (int, [][]int) {
    a := 0
    if image[sr][sc] != newColor {
        stack := []key{ key{sr, sc} }
        val := image[sr][sc]

        for len(stack) != 0 {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if pop.x > 0 && image[pop.y][pop.x - 1] == val {
                stack = append(stack, key{ pop.y, pop.x - 1 })
            }
            if pop.x < len(image[0]) - 1 && image[pop.y][pop.x + 1] == val {
                stack = append(stack, key{ pop.y, pop.x + 1 })
            }

            if pop.y > 0 && image[pop.y - 1][pop.x] == val {
                stack = append(stack, key{ pop.y - 1, pop.x })
            }
            if pop.y < len(image) - 1 && image[pop.y + 1][pop.x] == val {
                stack = append(stack, key{ pop.y + 1, pop.x })
            }

            if image[pop.y][pop.x] == 1 {
                a++
            }
            image[pop.y][pop.x] = newColor

        }
    }
    
    return  a ,image
}