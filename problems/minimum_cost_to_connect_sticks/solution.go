
import "container/heap"
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
    
    h := &IntHeap{}

    for _, i := range sticks {
        heap.Push(h, i)
    }
    sum := 0
    for h.Len() > 1 {
        v1 := heap.Pop(h).(int)

        v2 := heap.Pop(h).(int)
        sum += v1 + v2
        heap.Push(h, v1 + v2)

    }

    return sum

}