func validTree(n int, edges [][]int) bool {
    if n ==1 {
        return true
    }
    m := map[int][]int{}
    for _, i := range edges {
        m[i[0]] = append(m[i[0]], i[1])
        m[i[1]] = append(m[i[1]], i[0])
    }
    v := map[int]bool{}
    g := find(v, map[[2]int]bool{}, m, 0)
    if len(v) != n {
        return false
    }
    return g
}
func find(visited map[int]bool, path map[[2]int]bool , m map[int][]int, cur int) bool {
    for _, i := range m[cur] {
        if !path[[2]int{i, cur}] && visited[i] {
            return false
        }
        if !visited[i] {
            path[[2]int{cur, i}] = true
            visited[i] = true
            if !find(visited, path, m, i) {
                return false
            }
            path[[2]int{cur, i}] = false
            visited[i] = false
        }
    }
    return true
}