func longestSquareStreak(nums []int) int {
    m := map[int]int{}
    sort.Ints(nums)
    max := -1
    for _, i := range nums {
        m[i]++
    }
    for _, i := range nums {
        a := float64(len(find(m, i))) 
        if a > 0 {
            max = int(math.Max(a + 1, float64(max)))
        }
        
    }

    return max
    


}

func find(m map[int]int, a int) []int {
    if m[a*a] == 0 {
        return []int{}
    }
    return append([]int{a}, find(m, a*a)...)
}