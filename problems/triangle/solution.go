

func minimumTotal(triangle [][]int) int {
    for depth := len(triangle) - 2; depth >= 0; depth-- {
        for i := 0 ; i < len(triangle[depth]); i++ {

            triangle[depth][i] += int(math.Min(float64(triangle[depth +1][i]), float64(triangle[depth + 1][i + 1]))) 
        }
    }
    return triangle[0][0]
    
}
