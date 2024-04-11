func canCross(stones []int) bool {
    return re(0, stones, 0, map[[2]int]bool{})
}
func re(i int, stones []int, lastJump int, dp map[[2]int]bool) bool {

    if stones[len(stones) - 1] == i {
        dp[[2]int{i, lastJump}] = true
        return true
    }
    if val, ok := dp[[2]int{i, lastJump}]; ok {
        return val
    }
    if v := sort.SearchInts(stones, i); v >= len(stones) || stones[v] != i {
        dp[[2]int{i, lastJump}] = false
        return false
    }
    if lastJump > 0 && re(i + lastJump, stones, lastJump, dp) {
        dp[[2]int{i, lastJump}] = true
        return true
    }
    v1 := re(i + lastJump + 1, stones, lastJump + 1, dp)
    if v1 {
        dp[[2]int{i, lastJump}] = true
        return true
    }
    if lastJump > 1 {
        if re(i + lastJump - 1, stones, lastJump - 1, dp) {
            dp[[2]int{i, lastJump}] = true
            return true
        }
    }
    dp[[2]int{i, lastJump}] = false
    return false
}
