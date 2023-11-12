var visited = map[[3]int]bool{}

func distributeCandies(n int, limit int) int {
    visited = map[[3]int]bool{}

    return re(limit, limit, limit, n)
}
func re(a, b, c int, n int) int {
    if visited[[3]int{a, b, c}] {
        return 0
    }
    if a < 0 || b < 0 || c < 0 {
        return 0
    }
    visited[[3]int{a, b, c}] = true
    if n == 0 {

        return 1
    }
    
    return re(a -1, b, c, n - 1) + re(a, b - 1, c, n - 1) + re(a, b , c - 1, n - 1)
}