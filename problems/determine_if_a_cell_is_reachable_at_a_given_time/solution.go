func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    if sx == fx && sy == fy {
        return t != 1
    }
    xdif := abs(sx - fx)
    ydiff := abs(sy - fy)
    return max(xdif, ydiff) <= t
}
func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}
func max( i, j int) int {
    if i > j {
        return i
    }
    return j
}