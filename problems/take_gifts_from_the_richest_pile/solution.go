type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func pickGifts(gifts []int, k int) int64 {
    h := &IntHeap{}
    for _, i := range gifts {
        heap.Push(h, i)
    }
    for i := 0; i < k; i++ {
        pop := heap.Pop(h).(int)
        heap.Push(h, int(math.Sqrt(float64(pop))))
    }
    sum := int64(0)
    for h.Len() > 0 {
        sum += int64(heap.Pop(h).(int))
    }
    return sum
}