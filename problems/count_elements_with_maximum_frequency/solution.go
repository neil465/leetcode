func maxFrequencyElements(nums []int) int {
    m := map[int]int{}
    max := -1
    for _, v := range nums {
        m[v] ++
        if m[v] > max {
            max = m[v]
        }
    }

    k := 0
    for _, v := range nums {
        if m[v] == max {
            k ++
        }
    }
    return k
}