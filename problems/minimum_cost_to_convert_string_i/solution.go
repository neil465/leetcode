type qElm struct {
    curCost int
    value byte
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    move := map[string]int{}
    
    graph := map[byte][]qElm{}
    
    for i := range original {
        val := qElm{curCost : cost[i], value :changed[i]}
        graph[original[i]] = append(graph[original[i]], val)
    }
    
    for _, v := range original {
        queue := []qElm{qElm{curCost : 0, value : v}}
        
        for len(queue) > 0 {
            pop := queue[0]
            queue = queue[1:]
            if val, ok := move[string(v) + string(pop.value)]; ok {
                if pop.curCost >= val {
                    continue
                } 
            }
            
            move[string(v) + string(pop.value)] = pop.curCost
            
            for _, nextv := range graph[pop.value] {
                queue = append(queue, qElm{curCost : pop.curCost +nextv.curCost, value : nextv.value })
            }
        }
    }
    
    res := 0
    
    for i := range source {
        
        if source[i]== target[i]{
            continue
        }
        if val, ok := move[string(source[i]) + string(target[i])]; ok {

            res += val
        } else {
            return -1
        }
    }
    return int64(res)
}