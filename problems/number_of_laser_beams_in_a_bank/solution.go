func numberOfBeams(bank []string) int {
    c := []int{}
    for i := range bank {
        if bank[i] != strings.Repeat("0", len(bank[0])) {
            curCount := 0 

            for _, v := range bank[i] {
                if v == '1' {
                    curCount ++
                }
            }

            c = append(c, curCount)
        }
    }
    res := 0

    for i := 1 ;i < len(c); i++ {
        res += c[i - 1] * c[i]
    }
    return res
}