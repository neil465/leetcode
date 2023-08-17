func findMaxLength(nums []int) int {
    v := map[int]int{0 : -1}
    c := 0
    max := 0
    for i := range nums {
        if nums[i] == 1 {
            c ++
        } else {
            c --
        }
        if _, ok := v[c]; ok {
            // fmt.Println(v, i, v[c])
            max = int(math.Max(float64(max), float64(i - v[c])))
        } else {
            v[c] = i
        }
    }
    return max
}