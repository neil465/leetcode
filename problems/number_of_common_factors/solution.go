func commonFactors(a int, b int) int {
    sum := 0 
    for i := 1 ; i <= int(math.Min(float64(a),float64(b))) ; i++ {
        if a % i == 0 && b % i == 0 {
            sum++
        }
    }
    return sum
}