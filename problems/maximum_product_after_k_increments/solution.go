func maximumProduct(nums []int, k int) int {
    h := &IntHeap{}
    heap.Init(h)
    for _,i := range nums {
        heap.Push(h, i )
    }
    for i := 0 ; i < k ; i++ {
        min := heap.Pop(h).(int) + 1
        heap.Push(h, min)
    }
    sum := 1 
    for h.Len() > 0 {
        sum = (sum * heap.Pop(h).(int)) % 1000000007
    }
    return sum 
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
