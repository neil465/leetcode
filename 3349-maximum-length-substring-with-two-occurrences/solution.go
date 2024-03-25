func maximumLengthSubstring(s string) int {
    best := 0
    for i := range s {
        for j := i + 1; j <= len(s); j ++ {
            cnt := map[byte]int{}
            good := true
            for k := i ; k < j ; k++ {
                cnt[s[k]] ++ 
                if cnt[s[k]] > 2 {
                    good = false
                    break
                }
            }

            if good {
                best = max(best, j - i )
            }
        }
    }

    return best
}
