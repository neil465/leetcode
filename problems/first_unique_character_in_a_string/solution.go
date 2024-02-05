
import "math"
func firstUniqChar(s string) int {
    m := map[byte]int{}
    ind := map[byte]int{}
    for i := range s {
        m[s[i]] ++
        ind[s[i]] = i
    }
    minimum := math.MaxInt32
    for i := 0 ; i < 26 ;i ++ {
        if val, ok := m[byte('a' + i)]; ok && val == 1 {
            minimum = min(ind[byte('a' + i)], minimum)
        }
    }
    if minimum != math.MaxInt32 {
        return minimum
    }
    return -1
}