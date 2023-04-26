func splitNum(num int) int {
    digits := []int{}
    for num > 0  {
        digits = append(digits, num % 10)
        num /= 10
    }
    sort.Ints(digits)
    sum := 0
    multi := 1
    fmt.Println(digits)
    for i := len(digits) - 1; i >= 0 ; i -= 2 {
        sum += digits[i] * multi
        if i > 0 {
            sum += digits[i - 1] * multi
        }
        multi *= 10
    }
    return sum
}