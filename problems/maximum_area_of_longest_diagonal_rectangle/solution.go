func areaOfMaxDiagonal(dimensions [][]int) int {
    area, diag := -1, -1.0
    
    for _, i := range dimensions {
        v := math.Sqrt(float64(i[0] * i[0] + i[1] * i[1]))
        if v > diag {
            area = i[0] * i[1]
            diag = v
        }
        if v == diag {
            if area < i[0] * i[1] {
                area = i[0] * i[1]
            }
        }
    }
    return area
}