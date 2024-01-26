
import "math"

var dp = map[[2]int]int{}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    dp = map[[2]int]int{}
    connections := map[int][][]int{}

    for _, i := range flights {
        connections[i[0]] = append(connections[i[0]],i[1:])
    }
    v := r(connections, src, dst, k)
    if v == math.MaxInt32 {
        return -1
    }
    return v
    
}
func r (connections map[int][][]int, src, dst, k int) int {
    if k == -2 {
        return math.MaxInt32 
    }
    if src == dst {
        return 0
    }
    if val, ok := dp[[2]int{src, k}]; ok {
        return val
    }

    m := math.MaxInt32 
    
    for _, connection := range connections[src] {
        m = min(m, r(connections, connection[0], dst, k - 1) + connection[1])
    }
    dp[[2]int{src, k}] = m
    return m
}