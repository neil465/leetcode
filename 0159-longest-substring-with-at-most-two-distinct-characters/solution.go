func lengthOfLongestSubstringTwoDistinct(s string) int {
    m := map[byte]int{}
    i := 0
    best := 0
    for j := range s {

        

        if len(m) > 2 {
            mid := byte(" "[0])

            for l := range m {
                if m[l] < m[mid] || mid == byte(" "[0]) {
                    mid = l
                }
            }
            i = m[mid] + 1
            delete(m, mid)
        } 
        
        m[s[j]] = j

        best = max(best, j - i)
    }
    if len(m) > 2 {
        mid := byte(" "[0])

        for l := range m {
            if m[l] < m[mid] || mid == byte(" "[0]) {
                mid = l
            }
        }
        i = m[mid] + 1
        delete(m, mid)
    } 
    

    best = max(best, len(s) - i)
    return best

}
