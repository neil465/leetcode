func imageSmoother(img [][]int) [][]int {
    result := make([][]int, len(img))
    
    for i := 0; i < len(result); i++ {
        result[i] = make([]int, len(img[i]))
        copy(result[i], img[i])
    }
    
    for i := 0 ; i < len(img) ; i++{
        for j := 0 ; j < len(img[0]); j++{
            sum, count := 0,0
            for x := -1 ; x < 2 ; x ++{
                for y := -1 ; y < 2 ; y++{
                    s, c := requestCell(img, i + x, j + y)
                    sum += s
                    count += c
                }
            }
            result[i][j] = sum/count
        }
    }
    return result
}
func requestCell(b [][]int, x, y int) (int, int) {
    if x > len(b) - 1 || x < 0 || y > len(b[0]) - 1 || y < 0{
        return 0, 0
    }
    return b[x][y], 1
}