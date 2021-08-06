func firstMissingPositive(nums []int) int {
    m := make(map[int]bool)
    for _,i := range nums{
        m[i] = true
    }
    for i := 1; i < 1000000 ; i++{
        if m[i] != true{
            return i
        }
    }
    return 0
}