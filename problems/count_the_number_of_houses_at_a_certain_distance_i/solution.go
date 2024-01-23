
type qe struct {
    head, curval int
}

func countOfPairs(n int, x int, y int) []int {
    graph := map[int][]int{}

    visited  := map[[2]int]bool{}

    for i := 1; i < n; i ++ {
        graph[i] = append(graph[i], i + 1)
        graph[i + 1] = append(graph[i + 1], i)
    }

    graph[x] = append(graph[x], y)

    graph[y] = append(graph[y], x)

    queue := []qe{}

    for i := 1; i <= n ; i++ {
        visited[[2]int{i, i}] = true
        queue = append(queue, qe{i, i})
    }

    res := make([]int, n + 1)
    

    for i := 0; len(queue) > 0 ; i++{
        res[i] = len(queue)
        nqueue := []qe{}
        for _, v := range queue {
            for _, pos := range graph[v.curval] {
                if !visited[[2]int{v.head, pos}] {
                    visited[[2]int{v.head, pos}] = true
                    nqueue = append(nqueue, qe{v.head, pos})
                }
            }
        }

        queue = nqueue
    }

    return res[1:]
}