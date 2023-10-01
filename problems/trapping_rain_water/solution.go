func trap(height []int) int {
    l, r := make([]int, len(height)), make([]int, len(height))
    c1, c2 := 0, 0
    for i := 1; i < len(height); i++ {
        l[i], c1 = max(c1, height[i - 1]), max(c1, height[i - 1])
        r[len(height) - 1 - i], c2 = max(c2, height[len(height) - i]), max(c2, height[len(height) - i])
    }
    k := 0
    for i := range height {
        
        if min(l[i], r[i]) - height[i] <= 0 {
            // fmt.Println("pass")
            continue
        }
        // fmt.Println(l[i], r[i] , height[i], min(l[i], r[i]) - height[i])
        k += min(l[i], r[i]) - height[i]
    }
    return k
}
func max(i,j int)int {
    if i > j{
        return i
    }
    return j
}
func min(i,j int) int {
    if i < j {
        return i
    }
    return j
}