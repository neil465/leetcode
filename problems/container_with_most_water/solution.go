
func maxArea(height []int) int {
    max := 0 
    l, r := 0, len(height) - 1
    for l <= r {
        if max < int(math.Min(float64(height[l]), float64(height[r]))) * (r - l) {
            max = int(math.Min(float64(height[l]), float64(height[r]))) * (r - l)
        }
        if height[l] < height[r] {
            l++
        } else {
            r --
        }
    }
    return max
}
