func divisibilityArray(word string, m int) []int {
    remainder := 0
    res := []int{}
    for _, i := range word {
        digit := int(i) - 48
        remainder = ((remainder * 10) + digit) % m
        if remainder == 0 {
            res = append(res, 1)
        } else {
            res = append(res, 0)
        }
    }
    return res
}