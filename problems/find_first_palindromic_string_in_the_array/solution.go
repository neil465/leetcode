func firstPalindrome(words []string) string {

    for _, w := range words {
        i,j := 0, len(w) -1 
        k := true
        for i < j {
            if w[i] != w[j] {
                k = false
                break
            }
            i ++
            j --
        }
        if k {
            return w
        }
    }
    return ""
}