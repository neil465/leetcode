func isPossibleToSplit(nums []int) bool {
    m := map[int]int{}
    for _, i := range nums {
        m[i] ++
        if m[i] > 2 {
            return false
        }
    }
    return true
}