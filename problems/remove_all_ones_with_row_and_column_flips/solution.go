func removeOnes(grid [][]int) bool {
    row  := grid[0]
    invertrow := make([]int, len(row))
    copy(invertrow, row)
    for i := range invertrow{
        invertrow[i] = 1-invertrow[i]
    }
    fmt.Println(row)
    fmt.Println(invertrow)
    for _,i := range grid {
        if !reflect.DeepEqual(i, row) && !reflect.DeepEqual(i, invertrow){
            return false
        }
    }
    return true
    
}