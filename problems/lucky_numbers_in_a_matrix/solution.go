func luckyNumbers (matrix [][]int) []int {
    res:= []int{}
    for i := range matrix{
        for j , i2 := range matrix[i]{
            minimum := 1023241234
            maximum := -1023241234
            for _ , i3 := range matrix[i]{
                if i3<minimum{
                    minimum = i3
                } 
            }
            for _ , i3 := range matrix{
                if i3[j]>maximum{
                    maximum = i3[j]
                } 
            }
            if i2 == minimum && i2 == maximum{
                res = append(res,i2)
            }
        }
    }
    return res
}