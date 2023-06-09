func findDifferentBinaryString(nums []string) string {
    m := map[string]bool{}
    for _, i := range nums {
        m[i] = true
    }
    return find(m, "", len(nums[0]))
}
func find(m map[string]bool, s string, n int) string {
    if len(s) == n {
        if !m[s] { return s }
        return ""
    }
    a, b := find(m, s + "1", n), find(m, s + "0", n)
    if a != "" {
        return a 
    } else {
        return b
    }
}