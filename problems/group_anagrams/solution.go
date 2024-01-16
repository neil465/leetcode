func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}

    for _, i := range strs {
        a := [26]int{}

        for _, j := range i {
            a[int(j) - int('a')]++
        }
        m[a] = append(m[a], i)
    }
    k := [][]string{}

    for _, v := range m {
        k = append(k, v)
    }
    return k
}