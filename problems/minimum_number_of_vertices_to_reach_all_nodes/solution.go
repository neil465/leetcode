func findSmallestSetOfVertices(n int, edges [][]int) []int {
    m := map[int]bool{}

    for _, v := range edges {
        m[v[1]] = true
    }

    q := []int{}

    for i := 0 ; i < n ; i++ {
        if !m[i] {
            q = append(q, i)
        }
    }
    return q

}