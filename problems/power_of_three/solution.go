func isPowerOfThree(n int) bool {
    sum := 1 
    for sum < n {
        sum *= 3
    }
    return sum == n
}