func countKeyChanges(s string) int {
    k := 0
    for i := 0 ; i < len(s) - 1; i++ {
        if strings.ToLower(string(s[i])) != strings.ToLower(string(s[i + 1])) {
            k ++
        }
        
    }
    return k
}