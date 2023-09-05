func countSymmetricIntegers(low int, high int) int {
    res := 0 
    for i := low ; i <= high ; i++ {
        v := i 
        if (int(math.Log10(float64(v))) + 1) % 2 != 0 {
            continue
        }
        iter := (int(math.Log10(float64(v))) + 1) / 2
        firstSum := 0 
        nextSum := 0
        for j := 0 ; j < iter ; j ++ {
            firstSum += v % 10 
            v /= 10
        }
        for v > 0{
            nextSum += v % 10 
            v /= 10
        }
        if firstSum - nextSum == 0 {
            res ++
        }
    }
    return res
}