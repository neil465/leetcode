func findAndReplacePattern(words []string, pattern string) []string {
    value := ""
    m := map[rune]int{}
    j := 1
    for _,i := range pattern {
        if m[i] > 0 {
            value += fmt.Sprint(m[i])
        }else {
            j++
            value += fmt.Sprint(j)
            m[i] = j
        }
    }
    a := []string{}
    for _,k := range words {
        nv := ""
        m = map[rune]int{}
        j = 1
        for _,i := range k {
            if m[i] > 0 {
                nv += fmt.Sprint(m[i])
            }else {
                j++
                nv += fmt.Sprint(j)
                m[i] = j
                
            }
        }

        if nv == value {
            a = append(a, k)
        }
    }
    return a
}