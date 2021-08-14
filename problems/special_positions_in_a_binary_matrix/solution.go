func numSpecial(mat [][]int) int {
    row := make(map[int]int)
    collum := make(map[int]int)
    for o,i := range mat{
        for k,j := range i{
            if j == 1{
                row[o]++
                collum[k]++
            }
            
            
        }
    }
    sum := 0
    fmt.Println(row,collum)
    for o,i := range mat{
        for k,j := range i{
            
            if j == 1 && row[o] < 2 && collum[k] < 2{sum++}
        }
    }
    return sum
}