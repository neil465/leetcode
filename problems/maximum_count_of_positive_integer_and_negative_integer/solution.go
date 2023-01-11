func maximumCount(nums []int) int {
    neg, pos := 0, 0

    for _, i := range nums {
        if i < 0 {
            neg++
        } else if i > 0 {
            pos++
        }
    }

    return int(math.Max(float64(neg), float64(pos)))
}