func multiply(mat1 [][]int, mat2 [][]int) [][]int {
    var result = [][]int{}
    
    for _,row1 := range mat1 {
        var temp = []int{}
        
        for count := 0 ; count < len(mat2[0]) ; count++ {
            var sum int
            for pos := range mat2 {
                sum += row1[pos] * mat2[pos][count]
            }
            
            temp = append(temp,sum)
        }
        
        result = append(result,temp)
        
    }
    return result
}