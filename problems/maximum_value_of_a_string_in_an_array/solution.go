func maximumValue(strs []string) int {
    max := -1
    for _, i := range strs {
        val, err := strconv.Atoi(i)
        if err != nil && len(i) > max {
            max = len(i)
        }
        if err == nil && val > max{
            max = val
        }
    }
    return max
}