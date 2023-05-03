type Node struct {
    start int
    end int
    cost int
}

type Heap []Node

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func minCostConnectPoints(points [][]int) int {
    h := &Heap{}
    heap.Init(h)


    val := 0
    curNode := 0
    curEdges := 0
    v := map[int]bool{}
    for index := range points {
        heap.Push(h, Node{curNode, index, manhattanDistance(points[index][0], points[index][1], points[curNode][0], points[curNode][1])})
    }
    for {
        tempNode := heap.Pop(h).(Node)
        if !v[tempNode.end] {
            v[tempNode.end] = true
            val += tempNode.cost
            curEdges ++
            curNode = tempNode.end
            break
        }
    }
    for h.Len() > 0 && curEdges < len(points) {

        for index := range points {
            
            if index != curNode {
               
                heap.Push(h, Node{curNode, index, manhattanDistance(points[index][0], points[index][1], points[curNode][0], points[curNode][1])})
            // fmt.Println(abs(points[index][0] - points[curNode][0]) + abs(points[index][1] - points[curNode][1]))
            }
            
        }

        
        for {
            tempNode := heap.Pop(h).(Node)
            if !v[tempNode.end] {
                v[tempNode.end] = true
                val += tempNode.cost
                curEdges ++
                curNode = tempNode.end
                fmt.Println(tempNode.cost, points[tempNode.start], points[tempNode.end] )
                break
            }
            
        }
    }

    if curEdges != len(points) {
        return -1
    }
    return val
}
func manhattanDistance(x1, y1, x2, y2 int) int {
    return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func abs( i int) int {
    if i > 0 {
        return i
    }
    return i * -1
}