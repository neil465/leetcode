func isThree(n int) bool {
    g := int(math.Sqrt(float64(n)))*int(math.Sqrt(float64(n)))
    if g != n{return false}
    return IsPrime(int(math.Sqrt(float64(n))))
}
func IsPrime(v int) bool {
    for i := 2; i <= int(math.Floor(float64(v) / 2)); i++ {
        if v%i == 0 {
            return false
        }
    }
    return v > 1
}
