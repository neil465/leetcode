func findClosestNumber(nums []int) int {
    m := 2147483647
    for _, i := range nums {
        if abs(i) < abs(m) {
            m = i
        } else if abs(i) == abs(m) && i > m {
            m = i
        }
    }
    return m
}

func abs( i int) int {
    if i < 0 {
        return -i
    }
    return i
}