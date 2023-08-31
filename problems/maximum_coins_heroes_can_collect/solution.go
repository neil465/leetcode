type mon struct {
    c, p int
}

func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
    m := map[int][]int{}
    res := make([]int64, len(heroes))
    ms := []mon{}
    for i := range heroes { m[heroes[i]] = append(m[heroes[i]] , i) }
    for i := range monsters { ms = append(ms, mon{coins[i], monsters[i] }) }

    sort.Ints(heroes)
    sort.Slice(ms, func(i,j int) bool { return ms[i].p < ms[j].p })

    ind := 0 
    sum := 0
    for _, i := range heroes {
        for ind = ind ; ind < len(ms) && i >= ms[ind].p; ind ++ { sum += ms[ind].c }
        res[m[i][0]] = int64(sum)
        m[i] = m[i][1:]
    }
    return res
}