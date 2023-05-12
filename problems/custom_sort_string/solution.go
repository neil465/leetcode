func customSortString(order string, s string) string {
    lorder := [26]int{}
    for i := range order {
        lorder[int(order[i]) - 97] = i
    }
    a := []rune(s)
    sort.Slice(a, func(i, j int) bool {
        return lorder[int(a[i]) - 97] < lorder[int(a[j]) - 97]
    })
    return string(a)
}