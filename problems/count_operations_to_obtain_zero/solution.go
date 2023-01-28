func countOperations(num1 int, num2 int) int {
    if num1 == 0 || num2 == 0 {
        return 0
    }
    if num1 > num2 {
        return 1 + countOperations(num1 - num2, num2)
    }
    return 1 + countOperations(num1, num2 - num1)
}