/**
 * Definition for an infinite stream.
 * type InfiniteStream interface {
 *     Next() int
 * }
 */
func findPattern(stream InfiniteStream, pattern []int) int {
    lps := make([]int, len(pattern))
    j := 0 
    for i := 1 ; i < len(pattern); i ++ {
        for pattern[j] != pattern[i] && j > 0 {
            j = lps[j - 1]
        }

        if pattern[j] == pattern[i] {
            j ++
            lps[i] = j
        }
    }


    j = 0

    cur := []int{stream.Next()}

    for i := 0 ; i < len(cur); i ++ {
        if i + len(pattern) > len(cur) {
            i --
        } else if pattern[j] == cur[i] {
            if j == len(pattern) -1 {
                return i- len(pattern) + 1
            }
            
            j ++
        } else if j != 0 {
            j = lps[j - 1]
            i --
        }
       
        cur = append(cur, stream.Next())
    }
    return 0

}