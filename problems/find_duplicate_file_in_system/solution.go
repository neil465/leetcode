func findDuplicate(paths []string) [][]string {
    res := [][]string{}
    m := make(map[string][]string)
    
    for _,i := range paths{
        g := strings.Fields(i)
        for i := 1 ; i < len(g) ; i++{
            a := strings.Split(g[i],"(")
            c := m[a[1]]
            c = append(c,g[0]+"/"+a[0])
            m[a[1]] = c
        }
    }
    for _,i := range m{
        if len(i) > 1{
            res = append(res,i)
        }
    }
    return res
}