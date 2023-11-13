func findHighAccessEmployees(at [][]string) []string {
    sort.Slice(at, func(i, j int) bool{
        if at[i][0] != at[j][0] {
            return at[i][0] < at[j][0]
        }
        return at[i][1] < at[j][1]
    })

    a := map[string]bool{}
    for i := 0 ; i < len(at) - 2 ; i ++ {
        if at[i][0] == at[i + 2][0] && diffcheck(at[i][1], at[i + 2][1]) {
            a[at[i][0]] = true
        }
    }
    res := []string{}
    for i := range a {
        res = append(res, i)
    }
    return res
}

func diffcheck(s1, s2 string) bool {
    a1, _ := strconv.Atoi(s1)
    a2, _ := strconv.Atoi(s2)
    if a2 - a1 < 100 {
        return true
    }
    return false
}