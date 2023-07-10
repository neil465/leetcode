func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
    m := map[int]bool{}
    for _, i := range nums {
        m[i] = true
    }
    for i := range moveFrom {
        temp := m[moveFrom[i]]
        m[moveFrom[i]] = false
        m[moveTo[i]] = temp
        
    }

    a := []int{}
    for k, v := range m {
        if !v  {
            continue
        }
        a = append(a, k)
    }
    sort.Ints(a)
    return a
    
}