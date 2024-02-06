
import (
	"fmt"
	"math"
)

var dp = map[[26]int]int{}

func minStickers(stickers []string, target string) int {
    dp = map[[26]int]int{}
    m := map[string][26]int{}

    t := [26]int{}

    for _, val := range target {
        t[int(val ) - 'a'] ++
    }
    for _, sticker := range stickers {
        cur := [26]int{}
        for _, val := range sticker {
            if t[int(val) - 'a'] > cur[int(val )- 'a'] {
                cur[int(val )- 'a'] ++
            }
        }

        m[sticker] = cur
    }

    v := re(t, m, stickers) 
    if v == math.MaxInt32 {
        return -1
    }
    return v
}
func re(cur [26]int, m map[string][26]int, arr []string) int {
    
    b := true
    if val, ok := dp[cur]; ok {
        return val
    }
    for _, v := range cur {
        if v > 0 {
            b = false
            break
        }
    }
    if b  {
        return 0
    }
    minimum := math.MaxInt32
    branchingFactor := 0
    for _, sticker := range arr {
        newDiff := [26]int{}
        k := false
        for i := 0 ; i < 26; i++ {
            d := cur[i] - m[sticker][i]
            if d < 0 {
                d = 0
            }
            if d >=  0 && cur[i] != 0 &&  m[sticker][i] != 0 {
                k = true
            }
            newDiff[i] = d
        }
        if !k {
            continue
        }
        branchingFactor ++
        
        minimum = min(minimum, re(newDiff, m, arr) + 1)
    }

    
    dp[cur] = minimum
    return minimum
    
}
