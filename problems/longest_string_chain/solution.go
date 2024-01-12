
import "fmt"
func longestStrChain(words []string) int {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    fmt.Println(words)
    dp := make([]int, len(words))
    m := 0
    for i := len(words) -1 ; i >= 0 ; i -- {
        k := 0 
        for j := i  ; j < len(words); j ++ {
            if len(words[j]) - len(words[i]) == 1 {
                icur, jcur := 0,0
                b := true
                bad := false
                for icur < len(words[i]) {
                    if words[i][icur] != words[j][jcur] {
                        if b {
                            icur --
                            b = false
                        } else {
                            bad = true

                            break
                        }
                    }
                    icur ++
                    jcur ++
                }
                if !bad {
                    k = max(k, dp[j] + 1)
                }
            }
        }
        dp[i] = k
        m = max(dp[i], m)
    }
    fmt.Println(dp)
    return m + 1
}