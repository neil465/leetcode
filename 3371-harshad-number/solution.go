func sumOfTheDigitsOfHarshadNumber(x int) int {
    dsum := 0

    v := x 

    for v > 0 {
        dsum += v % 10
        v /= 10
    }

    if x % dsum == 0 {
        return dsum
    }
    return -1
}
