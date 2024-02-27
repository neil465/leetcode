
func treeDiameter(edges [][]int) int {
    g := map[int][]int{}

    for _, v := range edges {
        g[v[0]] = append(g[v[0]], v[1])
        g[v[1]] = append(g[v[1]], v[0])
    }
    return diameterOfBinaryTree(0, g)
}

var best = 0
func diameterOfBinaryTree(root int, g map[int][]int) int {
    best = 0
    re(root, g, map[int]bool{})
    return best
}
func re(cur int, g map[int][]int, visited map[int]bool) int {
    if visited[cur] {
        return 0
    }
    arr := []int{}
    for _, nextNode := range g[cur] {
        visited[cur] = true
        arr = append(arr, re(nextNode, g, visited) )
        visited[cur] = false
    }
    if len(arr) == 0 {
        return 0
    }
    if len(arr) == 1 {
        best = max(best, arr[0])
        return arr[0] + 1
    }
    sort.Ints(arr)
    v1, v2 := arr[len(arr) - 1], arr[len(arr) - 2]
    best = max(v1 + v2 , best)
    return max(v1, v2) + 1
}