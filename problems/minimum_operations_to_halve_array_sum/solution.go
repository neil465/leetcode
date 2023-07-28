type IntHeap []float64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func halveArray(nums []int) int {
    s := 0.0
    h := &IntHeap{}
    heap.Init(h)
    for _, i := range nums {
        s += float64(i)
        heap.Push(h, float64(i))
    }
    k := 0.0
    g := 0
    for k * 2 < s {
        a := heap.Pop(h).(float64)
        k += a / 2.0
        heap.Push(h, a / 2.0)
        g ++
    }
    return g
}