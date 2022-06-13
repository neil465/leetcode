func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, j := range nums {
        if m[target - j] > 0 {
            return []int{m[target - j] - 1, i}
        }
        m[j] = i + 1
    }
    return []int{}
}