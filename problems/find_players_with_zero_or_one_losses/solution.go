func findWinners(matches [][]int) [][]int {
    m := map[int]int{}

    for _, i := range matches {
        m[i[1]] ++
        m[i[0]] = m[i[0]]
    }

    k := [][]int{{}, {}}

    for key, v := range m {
        if v == 0 {
            k[0] = append(k[0], key)
        } else if v == 1 {
            k[1] = append(k[1], key)
        }
    }
    sort.Ints(k[0])
    sort.Ints(k[1])
    return k
}