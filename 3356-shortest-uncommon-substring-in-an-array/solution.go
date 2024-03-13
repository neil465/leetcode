func shortestSubstrings(arr []string) []string {
    m := map[string]int{}

    for _, v := range arr {
        curmap := map[string]bool{}
        for i := range v {
            for j := i; j < len(v); j ++ {
                if curmap[v[i:j +1]] {
                    continue
                }
                m[v[i:j +1]] ++
                curmap[v[i:j +1]] = true
            }
        }
    }
    res := []string{}
    for _, v := range arr {
        best := ""
        for i := range v {
            for j := i; j < len(v); j ++ {
                if m[v[i:j +1]] == 1 {
                    if best == ""{
                        best = v[i:j +1]
                    }
                    best = lexMin(best, v[i:j +1])
                }
            }
        }
        res = append(res, best)
    }
    return res
}

func lexMin(s1, s2 string) string {
    if len(s1) < len(s2) {
        return s1
    } else if len(s1) > len(s2) {
        return s2
    }
    // If lengths are the same, compare lexicographically
    if s1 < s2 {
        return s1
    }
    return s2
}
