func titleToNumber(columnTitle string) int {
    sum := 0 
    for i := 0 ; i < len(columnTitle) ; i++ {
        sum += int(math.Pow(26, float64(len(columnTitle) - i) - 1) * float64(columnTitle[i] - 64))
    }
    return sum
}