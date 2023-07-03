type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func fillCups(amount []int) int {
    h := &IntHeap{}
    heap.Init(h)
    for _, i := range amount {
        heap.Push(h, i)
    }
    t := 0
    for {
        a, b := heap.Pop(h).(int), heap.Pop(h).(int)
        if a == 0 {
            break
        }
        heap.Push(h, a - 1)
        heap.Push(h, b - 1)
        t ++
    }
    return t
}