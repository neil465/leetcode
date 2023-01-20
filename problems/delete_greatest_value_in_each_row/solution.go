func deleteGreatestValue(grid [][]int) int {
    n := len(grid[0])
    sum := 0 
    for iteration := 0 ; iteration < n ; iteration ++ {
        a := 0
        for index := range grid {
            if iteration == 0 {
                sort.Ints(grid[index])
            }
            a = int(math.Max(float64(grid[index][len(grid[index]) - 1]), float64(a)))

            grid[index] = grid[index][:len(grid[index]) - 1]
        }
        
        fmt.Println(a, "b")
        sum += a
    }

    return sum
    
}