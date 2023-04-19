func smallestNumber(num int64) int64 {
    g := true
    if num < 0 {
        g = false
        num *= -1
    }
    digits := []int64{}
    z := 0
    for num > 0 {
        digits = append(digits, num%10)
        if digits[len(digits) -1 ] == 0 {
            z ++
            digits = digits[:len(digits) - 1]
        }
        num /= 10
    }
    sort.Slice(digits, func(i, j int ) bool {
        if g {
            return digits[i] < digits[j]
        }
        return digits[i] > digits[j]
    })
    m := 1
    sum := int64(0)
    if !g {
        for i := 0; i < z ; i++ {
            digits = append(digits, 0)
        }
    } else {
        for i := 0; i < z ; i++ {
            digits = append([]int64{digits[0]}, append([]int64{0}, digits[1:]... )...)
        }
    }
    for i := len(digits) - 1; i >= 0 ; i-- {
        sum += digits[i] * int64(m)
        m *= 10
    } 
    if !g {
        return sum * -1
    }
    return sum
}