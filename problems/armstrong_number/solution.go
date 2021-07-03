func isArmstrong(n int) bool {
    b := n
    res := 0 
    c := int(math.Log10(float64(n))+1)
    for n > 0 {
        
        res += int(math.Pow(float64(n%10),float64(c)))
        n/=10
    }
    return b == res
}