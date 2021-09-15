func frequencySort(nums []int) []int {
    m := make( map[int] []int )
    h := &IntHeap{}
    res := []int{}
    for _, num := range nums {
        m[num] = append(m[num], num)
    }
    
    for _, v := range m {
        heap.Push(h, v)
    }
    
    for h.Len() > 0 {
        res = append(res, heap.Pop(h).([]int)...)
    }
    
    return res
}

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { 
    if len(h[i]) == len(h[j]) {
        return h[i][0] > h[j][0]
    }
    return len(h[i]) < len(h[j]) 
}
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
