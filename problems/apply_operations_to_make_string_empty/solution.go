func lastNonEmptyString(s string) string {
    chars := [26][]int{}

    best := []int{}

    for ind, v := range s {
        chars[int(v - 'a')] = append(chars[int(v-'a')], ind)
        
    }

    for i, v := range chars {
        if len(best) == 0 {
            best = []int{i}
        } else if len(chars[i]) > len(chars[best[0]]) {
            best = []int{i}
        } else if len(best) == 0 || len(v) == len(chars[best[0]]) {
            best = append(best, i)
        }
    }
    res := [][]int{}

    for _, i := range best {
        res = append(res, []int{i, chars[i][len(chars[i]) - 1]})
    }

    sort.Slice(res, func(i, j int) bool {
        return res[i][1] < res[j][1]
    })

    res2 := ""

    for _, i := range res {
        res2 += string(i[0] + 'a')
    }
    return res2
}