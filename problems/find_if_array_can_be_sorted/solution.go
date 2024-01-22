
import (
	"fmt"
	"sort"
)

type s struct {
    v, i int
}
func canSortArray(nums []int) bool {
    m := map[int]int{}

    for _, num := range nums {
        count := 0 

        cur := 1

        for i := 0 ; i < 33; i++ {
            if num & cur > 0 {
                count++
            }
            cur = cur << 1
        }

        m[num] = count
    }
    fmt.Println(m)
    copyNums := make([]s, len(nums))

    for i := range nums {
        copyNums[i] = s{nums[i], i}
    }
    
    sort.Slice(copyNums, func(i, j int) bool {
        return copyNums[i].v < copyNums[j].v
    })

    for i := range copyNums {
        fmt.Println(copyNums[i], nums[i])
        cur := m[copyNums[i].v]
        for j := min(copyNums[i].i, i); j <= max(copyNums[i].i, i); j ++ {
            if m[nums[j]] != cur {
                return false
            }
        }
    }
    return true
}