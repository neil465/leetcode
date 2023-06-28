func validPath(n int, edges [][]int, source int, destination int) bool {
    m := map[int][]int{}
    for _, i := range edges {
        m[i[0]] = append(m[i[0]], i[1])
        m[i[1]] = append(m[i[1]], i[0])
    }
    a := []int{source}
    v := map[int]bool{}
    for len(a) > 0 {
        p := a[0]
        a = a[1:]
        if v[p] {
            continue
        }
        
        if p == destination {
            return true
        }
        v[p] = true
        a = append(a, m[p]...)
    }
    return false
}