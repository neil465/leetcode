func shortestBeautifulSubstring(s string, k int) string {
    bsubstr := ""
    i, j := 0, 0
    curcount := 0
    if s[0] == '1' {
        curcount ++
    }
    if curcount == k {
        bsubstr = "1"
    }
    
    for j < len(s) {
        // fmt.Println(bsubstr, s[i:j])
        if curcount >= k {
            if s[i] == '1' {
                curcount --
            }
            i ++ 
        } else if curcount < k {
            j ++
            if j < len(s) {
                if s[j] == '1' {
                    curcount ++
                }
            }
        }
        if curcount == k {
            if bsubstr == "" ||( len(bsubstr) > j - i && i <= j && isLexSmaller(s[i:j + 1], bsubstr)){
                
                bsubstr = s[i:j + 1]
                
            }
            
        }
    }
    return bsubstr
}
func isLexSmaller(l1, l2 string) bool {
    if len(l1) != len(l2){ 
        return min(len(l1), len(l2)) == len(l1)
    }
    for i := 0; i < min(len(l1), len(l2)); i ++ {
        if l1[i] < l2[i] {
            return true
        } else if l1[i] > l2[i] {
            return false
        }
    }
    return min(len(l1), len(l2)) == len(l1)
    
}
func min(i, j int)int {
    if i < j {
        return i
    }
    return j
}