func alternateDigitSum(n int) int {
    sum, sign := 0, 1
    if int(math.Log10(float64(n)) + 1) % 2 == 0{
        sign *= -1
    }
    for n > 0 {
        sum += sign * (n%10)
        fmt.Println("p1", n, sum)
        n /= 10
        fmt.Println("p2", n, sum)
        sign *= -1
    }

    return sum
}