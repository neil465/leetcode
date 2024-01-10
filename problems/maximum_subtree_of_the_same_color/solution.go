
var m = math.MinInt32

func maximumSubtreeSize(edges [][]int, colors []int) int {
    graph := map[int][]int{}
    m = math.MinInt32
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }

    a(graph, colors, 0, map[int]bool{})
    return m
}
func a(graph map[int][]int, colors []int, cur int, visited map[int]bool) int {
    value := 1
    bret := false
    for _, i := range graph[cur] {
        if visited[i]  {
            continue
        }
        
        
        visited[cur] = true
        v := a(graph, colors, i, visited)
        visited[cur] = false
        if colors[i] != colors[cur] {
            bret = true
        }
        if v == -1 {
            bret = true
        }
        value += v
    }
    if bret {
        return - 1
    }
    m = max(m, value )
    return value

}