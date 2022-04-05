func maxArea(height []int) int {
    max := 0 
    for i, j := 0, len(height) - 1; i <= j; i = i {
        max = compare(max, compare(height[i], height[j], false) * (j - i), true)
        if height[i] < height[j] {
            i++
        }else{
            j--
        }
    }
    return max
}
func compare(i, j int, b bool ) int {
    if b { if i > j { return i }; return j } else { if i > j { return j } else { return i } }
}