func cellsInRange(s string) []string {
    a := []string{}
    first, _ := strconv.Atoi(string(s[1]))
    last, _ := strconv.Atoi(string(s[4]))
    for i := int(s[0]); i < int(s[3]) + 1; i++{
        for j := first; j <= last; j++{
            a = append(a, string(i) + strconv.Itoa(j) )
        }
    }
    return a
}